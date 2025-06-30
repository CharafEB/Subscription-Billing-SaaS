package router

import (
	"github.com/go-chi/chi/v5"
)

func (app *GetWayController) subscription(r chi.Router) chi.Router {

	r.Post("/", app.SubscritionHandler.Subscribtion)

	return r
}
