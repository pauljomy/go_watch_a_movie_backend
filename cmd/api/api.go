package main

import (
	"log/slog"
	"net/http"
	"time"

	"backend/cmd/api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	logger  *slog.Logger
	config  config
	handler *handlers.Handler
}

type config struct {
	addr string
}

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(logRequest(app.logger))
	mux.Use(middleware.Recoverer)
	mux.Use(commonHeaders)
	mux.Use(enableCORS)

	mux.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.handler.Health)
		r.Get("/", app.handler.Home)
		r.Get("/movies", app.handler.GetAllMovies)
		r.Post("/authenticate", app.handler.Authenticate)
	})

	return mux
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	return srv.ListenAndServe()
}
