package router

import (
	"github.com/go-chi/chi/v5"
)

func (app *GetWayController) castmer(r chi.Router) chi.Router {

	r.Route("/Singin", func(r chi.Router) {
		r.Post("/", app.CastmerHandler.SingUp)
	})

	r.Route("/Login", func(r chi.Router) {
		r.Post("/", app.CastmerHandler.Login)
	})
	return r
}
