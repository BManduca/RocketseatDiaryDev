package api

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/BManduca/RocketseatDiaryDev/tree/main/EventWeekGoReact/internal/store/pgstore"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// chamando o método ServeHttp do meu router(handler)
	h.r.ServeHTTP(w, r)
}

// http.Handler -> é um tipo que consegue responder a um httpRequest
func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q: q,
	}

	r := chi.NewRouter()
	/*
		middleware -> atua como uma ponte entre diversas tecnologias,
		ferramentas e bancos de dados para integrá-los perfeitamente
		em um único sistema.
	*/
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

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

// inscrevendo um usuário em uma sala
func (h apiHandler) handleSubscribe(w http.ResponseWriter, r *http.Request) {}

// criando uma sala
func (h apiHandler) handleCreateRoom(w http.ResponseWriter, r *http.Request) {}

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
