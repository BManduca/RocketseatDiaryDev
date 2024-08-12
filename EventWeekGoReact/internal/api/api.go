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
	q *pgstore.Queries
	r *chi.Mux
	// basicamente o upgrader, serve para realizar um upgrade da nossa requisição HTTP rest, para se tornar uma requisição WebSocket
	upgrader websocket.Upgrader
	// map -> responsável por guardar todas as conexões em aberto com os clientes
	subscribers map[string]map[*websocket.Conn]context.CancelFunc
	mu          *sync.Mutex // mutual exclusion
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// chamando o método ServeHttp do meu router(handler)
	h.r.ServeHTTP(w, r)
}

// http.Handler -> é um tipo que consegue responder a um httpRequest
func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q: q,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		// inicialiando o map, pois, os maps em go, inicializam como null
		subscribers: make(map[string]map[*websocket.Conn]context.CancelFunc),
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

			// route messages
			r.Route("/{room_id}/messages", func(r chi.Router) {
				r.Post("/", a.handleCreateRoomMessage)
				r.Get("/", a.handleGetRoomMessages)

				r.Route("/{message_id}", func(r chi.Router) {
					r.Get("/", a.handleGetRoomMessage)
					r.Patch("/react", a.handleReactioToMessage)
					r.Delete("/react", a.handleRemoveReactioFromMessage)
					r.Patch("/answered", a.handleMarkMessageAsAnswered)
				})
			})
		})
	})

	a.r = r

	return a
}

/*
	vale ressaltar que mesmo trabalhando com WebSocket e apiRest,
	a assinatura da(o) função/método é a mesma
*/

// inscrevendo um usuário em uma sala | método que vai permitir que faça webSocket usando go
func (h apiHandler) handleSubscribe(w http.ResponseWriter, r *http.Request) {
	rawRoomID := chi.URLParam(r, "room_id")
	// verificando se é um ID válido
	roomID, err := uuid.Parse(rawRoomID)
	if err != nil {
		http.Error(w, "Invalid room id", http.StatusBadRequest)
		return
	}

	// 'buscando' la do DB e verificando se a sala existe
	// conexao com cliente pelo lado do servidor
	_, err = h.q.GetRoom(r.Context(), roomID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "room not found", http.StatusBadRequest)
			return
		}
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	c, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Warn("Failed to upgrade connection", "error", err)
		http.Error(w, "Failed to upgrade to ws connection", http.StatusBadRequest)
		return
	}

	// limpeza de recursos
	defer c.Close()

	ctx, cancel := context.WithCancel(r.Context())

	h.mu.Lock()
	// verificando se o map já existe ou ja foi inicializado
	if _, ok := h.subscribers[rawRoomID]; !ok {
		h.subscribers[rawRoomID] = make(map[*websocket.Conn]context.CancelFunc)
	}
	slog.Info("New client connected", "room_id", rawRoomID, "client_ip", r.RemoteAddr)
	h.subscribers[rawRoomID][c] = cancel //conexao cliente
	h.mu.Unlock()

	<-ctx.Done()

	h.mu.Lock()
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
		http.Error(w, "Invalid json", http.StatusBadRequest)
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

	data, _ := json.Marshal(response{ID: roomID.String()})
	// setando o header
	w.Header().Set("Content-Type", "application/json")
	// escrevendo os dados la no client
	_, _ = w.Write(data)
}

// retornar todas as salas
func (h apiHandler) handleGetRooms(w http.ResponseWriter, r *http.Request) {}

// escrevendo uma mensagem
func (h apiHandler) handleCreateRoomMessage(w http.ResponseWriter, r *http.Request) {}

// pegando todas as mensagens de uma sala especifica
func (h apiHandler) handleGetRoomMessages(w http.ResponseWriter, r *http.Request) {}

// pegando uma mensagem especifica de uma sala
func (h apiHandler) handleGetRoomMessage(w http.ResponseWriter, r *http.Request) {}

// reagindo a uma mensagem
func (h apiHandler) handleReactioToMessage(w http.ResponseWriter, r *http.Request) {}

// removendo a reação de uma mensagem
func (h apiHandler) handleRemoveReactioFromMessage(w http.ResponseWriter, r *http.Request) {}

// marcando uma mensagem como respondida
func (h apiHandler) handleMarkMessageAsAnswered(w http.ResponseWriter, r *http.Request) {}
