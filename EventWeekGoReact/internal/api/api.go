package api

import (
	"net/http"

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

	a.r = r

	return a
}
