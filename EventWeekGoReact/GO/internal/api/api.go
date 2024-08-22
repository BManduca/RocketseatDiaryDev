package api

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"sync"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5"

	"github.com/BManduca/RocketseatDiaryDev/tree/main/EventWeekGoReact/internal/store/pgstore"
)

type apiHandler struct {
	// criando a struct
	q           *pgstore.Queries
	r           *chi.Mux
	upgrader    websocket.Upgrader                                // basicamente o upgrader, serve para realizar um upgrade da nossa requisição HTTP rest, para se tornar uma requisição WebSocket
	subscribers map[string]map[*websocket.Conn]context.CancelFunc // map -> responsável por guardar todas as conexões em aberto com os clientes
	mu          *sync.Mutex                                       // mutual exclusion
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// chamando o método ServeHttp do meu router(handler)
	h.r.ServeHTTP(w, r)
}

// http.Handler -> é um tipo que consegue responder a um httpRequest
func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q:           q,
		upgrader:    websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }},
		subscribers: make(map[string]map[*websocket.Conn]context.CancelFunc), // inicialiando o map, pois, os maps em go, inicializam como null
		mu:          &sync.Mutex{},
	}

	r := chi.NewRouter()
	/*
		middleware -> atua como uma ponte entre diversas tecnologias,
		ferramentas e bancos de dados para integrá-los perfeitamente
		em um único sistema.
	*/
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/subscribe/{room_id}", a.handleSubscribe)

	r.Route("/api", func(r chi.Router) {
		// room routes
		r.Route("/rooms", func(r chi.Router) {
			r.Post("/", a.handleCreateRoom)
			r.Get("/", a.handleGetRooms)

			r.Route("/{room_id}", func(r chi.Router) {
				r.Get("/", a.handleGetRoom)

				// route messages
				r.Route("/messages", func(r chi.Router) {
					r.Post("/", a.handleCreateRoomMessage)
					r.Get("/", a.handleGetRoomMessages)

					r.Route("/{message_id}", func(r chi.Router) {
						r.Get("/", a.handleGetRoomMessage)
						r.Patch("/react", a.handleReactionToMessage)
						r.Delete("/react", a.handleRemoveReactionFromMessage)
						r.Patch("/answered", a.handleMarkMessageAsAnswered)
					})
				})
			})
		})
	})

	a.r = r

	return a
}

const (
	MessageKindMessageCreated           = "message_created"
	MessageKindReactionIncreasedMessage = "message_reaction_increased"
	MessageKindReactionDecreasedMessage = "message_reaction_decreased"
	MessageKindAnsweredMessage          = "message_answered"
)

type MessageReactionIncreasedMessage struct {
	ID    string `json:"id"`
	Count int64  `json:"count"`
}

type MessageReactionDecreasedMessage struct {
	ID    string `json:"id"`
	Count int64  `json:"count"`
}

type MessageAnsweredMessage struct {
	ID string `json:"id"`
}

type MessageCreatedMessage struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type Message struct {
	Kind   string `json:"kind"`  // tipo da message
	Value  any    `json:"value"` // value da message
	RoomID string `json:"-"`     // uso interno -> Vincular message com room
}

func (h apiHandler) notifyClients(msg Message) {
	h.mu.Lock()
	defer h.mu.Unlock()

	subscribers, ok := h.subscribers[msg.RoomID]
	if !ok || len(subscribers) == 0 {
		return
	}

	for conn, cancel := range subscribers {
		if err := conn.WriteJSON(msg); err != nil {
			slog.Error("Failed To Send Notify To Client", "error", err)
			cancel()
		}
	}
}

/*
	vale ressaltar que mesmo trabalhando com WebSocket e apiRest,
	a assinatura da(o) função/método é a mesma
*/

// inscrevendo um usuário em uma sala | método que vai permitir que faça webSocket usando go
func (h apiHandler) handleSubscribe(w http.ResponseWriter, r *http.Request) {
	_, rawRoomID, _, ok := h.readRoom(w, r)
	if !ok {
		return
	}

	c, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Warn("Failed to Upgrade Connection", "error", err)
		http.Error(w, "Failed to Upgrade to WebSocket Connection", http.StatusBadRequest)
		return
	}

	/*
		Para quando esse requerimento acabar
		iremos fechar a conexão com o client,
		ou seja, uma limpeza de recursos.
	*/
	defer c.Close()

	ctx, cancel := context.WithCancel(r.Context())

	h.mu.Lock() //travando o mutex para add uma conn nova
	// verificando se o map já existe ou ja foi inicializado
	if _, ok := h.subscribers[rawRoomID]; !ok {
		// se caso a sala não existir no meu map, estaremos criando a conexão
		h.subscribers[rawRoomID] = make(map[*websocket.Conn]context.CancelFunc)
	}
	slog.Info("New Client Connected", "room_id", rawRoomID, "client_ip", r.RemoteAddr)
	h.subscribers[rawRoomID][c] = cancel //se caso a sala já existir no map, a conexao é cancelada
	h.mu.Unlock()

	/*
		bloqueando a função com meu client, aguardando ate o context dar 'done'
		=> Duas formas de encerrar o processo, seja pelo cliente, quanto servidor.
	*/
	<-ctx.Done()

	// removendo o subscriber da minha lista de subscribers
	h.mu.Lock()
	// removendo conexao do meu pull de conexões
	delete(h.subscribers[rawRoomID], c)
	h.mu.Unlock()
}

// criando uma sala
func (h apiHandler) handleCreateRoom(w http.ResponseWriter, r *http.Request) {
	// body
	type _body struct {
		Theme string `json:"theme"`
	}

	var body _body
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	roomID, err := h.q.InsertRoom(r.Context(), body.Theme)
	if err != nil {
		slog.Error("Failed to insert room", "error", err)
		http.Error(w, "Sometthing went wrong", http.StatusInternalServerError)
		return
	}

	type response struct {
		ID string `json:"id"`
	}

	sendJSON(w, response{ID: roomID.String()})
}

// retornar todas as salas
func (h apiHandler) handleGetRooms(w http.ResponseWriter, r *http.Request) {
	// buscando no DB todas as salas existentes
	rooms, err := h.q.GetRooms(r.Context())
	if err != nil {
		slog.Error("Failed To Get Rooms", "error", err)
		http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
		return
	}

	if rooms == nil {
		rooms = []pgstore.Room{}
	}

	sendJSON(w, rooms)
}

func (h apiHandler) handleGetRoom(w http.ResponseWriter, r *http.Request) {
	room, _, _, ok := h.readRoom(w, r)

	if !ok {
		return
	}

	sendJSON(w, room)
}

// escrevendo uma mensagem
func (h apiHandler) handleCreateRoomMessage(w http.ResponseWriter, r *http.Request) {
	_, rawRoomID, roomID, ok := h.readRoom(w, r)

	if !ok {
		return
	}

	type _body struct {
		Message string `json:"message"`
	}
	var body _body
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	messageID, err := h.q.InsertMessage(r.Context(), pgstore.InsertMessageParams{
		RoomID:  roomID,
		Message: body.Message,
	})
	if err != nil {
		slog.Error("Failed To Insert Message", "error", err)
		http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
		return
	}

	type response struct {
		ID string `json:"id"`
	}

	sendJSON(w, response{ID: messageID.String()})

	// chamando a function Notify de maneira assíncrona
	go h.notifyClients(Message{
		Kind:   MessageKindMessageCreated,
		RoomID: rawRoomID,
		Value: MessageCreatedMessage{
			ID:      messageID.String(),
			Message: body.Message,
		},
	})
}

// pegando todas as mensagens de uma sala especifica
func (h apiHandler) handleGetRoomMessages(w http.ResponseWriter, r *http.Request) {
	_, _, roomID, ok := h.readRoom(w, r)

	if !ok {
		return
	}

	messages, err := h.q.GetRoomMessages(r.Context(), roomID)
	if err != nil {
		slog.Error("Failed To Get Room Messages", "error", err)
		http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
		return
	}

	if messages == nil {
		messages = []pgstore.Message{}
	}

	sendJSON(w, messages)
}

// pegando uma mensagem especifica de uma sala
func (h apiHandler) handleGetRoomMessage(w http.ResponseWriter, r *http.Request) {
	_, _, _, ok := h.readRoom(w, r)

	if !ok {
		return
	}

	rawMessageID := chi.URLParam(r, "message_id")
	messageID, err := uuid.Parse(rawMessageID)

	if err != nil {
		http.Error(w, "Invalid Message ID", http.StatusBadRequest)
	}

	message, err := h.q.GetMessage(r.Context(), messageID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Message Not Found", http.StatusBadRequest)
		}

		slog.Error("Failed To Get Message", "error", err)
		http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
		return
	}

	sendJSON(w, message)
}

// reagindo a uma mensagem
func (h apiHandler) handleReactionToMessage(w http.ResponseWriter, r *http.Request) {
	_, rawRoomID, _, ok := h.readRoom(w, r)

	if !ok {
		return
	}

	rawID := chi.URLParam(r, "message_id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		http.Error(w, "Invalid Message ID", http.StatusBadRequest)
		return
	}

	count, err := h.q.ReactToMessage(r.Context(), id)
	if err != nil {
		slog.Error("Failed To React To Message", "error", err)
		http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
		return
	}

	type response struct {
		Count int64 `json:"count"`
	}

	sendJSON(w, response{Count: count})

	go h.notifyClients(Message{
		Kind:   MessageKindReactionIncreasedMessage,
		RoomID: rawRoomID,
		Value: MessageReactionIncreasedMessage{
			ID:    rawID,
			Count: count,
		},
	})
}

// removendo a reação de uma mensagem
func (h apiHandler) handleRemoveReactionFromMessage(w http.ResponseWriter, r *http.Request) {
	_, rawRoomID, _, ok := h.readRoom(w, r)

	if !ok {
		return
	}

	rawID := chi.URLParam(r, "message_id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		http.Error(w, "Invalid Message ID", http.StatusBadRequest)
		return
	}

	count, err := h.q.RemoveReactionFromMessage(r.Context(), id)
	if err != nil {
		slog.Error("Failed To Decrease React To Message", "error", err)
		http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
		return
	}

	type response struct {
		Count int64 `json:"count"`
	}

	sendJSON(w, response{Count: count})

	go h.notifyClients(Message{
		Kind:   MessageKindReactionDecreasedMessage,
		RoomID: rawRoomID,
		Value: MessageReactionDecreasedMessage{
			ID:    rawID,
			Count: count,
		},
	})
}

// marcando uma mensagem como respondida
func (h apiHandler) handleMarkMessageAsAnswered(w http.ResponseWriter, r *http.Request) {
	_, rawRoomID, _, ok := h.readRoom(w, r)

	if !ok {
		return
	}

	rawID := chi.URLParam(r, "message_id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		http.Error(w, "Invalid Message ID", http.StatusBadRequest)
		return
	}

	err = h.q.MarkMessageAsAnswered(r.Context(), id)
	if err != nil {
		slog.Error("Failed To Mark Message As Answered", "error", err)
		http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	go h.notifyClients(Message{
		Kind:   MessageKindAnsweredMessage,
		RoomID: rawRoomID,
		Value: MessageAnsweredMessage{
			ID: rawID,
		},
	})
}
