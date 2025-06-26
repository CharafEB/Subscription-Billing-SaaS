package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *GetWayController) Castmer(r chi.Router) chi.Router {

	r.Route("/Singin", func(r chi.Router) {
		r.Post("/", app.CastmerHandler.SingUp)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Castmer Singin"))
		})
	})
	return r
}
