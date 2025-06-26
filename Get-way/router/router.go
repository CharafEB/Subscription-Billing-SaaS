package router

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (cnt *GetWayController) Mux() *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
	)
	r.Mount("/", cnt.Castmer(r))
	return r
}

func (app *Application) Run(mux *chi.Mux) error {
	if app.Address == "" {
		app.Address = ":8000"
	}

	var handler http.Handler = mux
	srv := &http.Server{
		Addr:         app.Address,
		Handler:      handler,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server is starting on Port %s", app.Address)

	if err := srv.ListenAndServe(); err != nil {
		log.Printf("Server failed to start: %v", err)
		return err
	}
	return nil
}
