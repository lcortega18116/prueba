package httpapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"

	"github.com/lcortega18116/prueba/internal/config"
	"github.com/lcortega18116/prueba/internal/handlers"
	middleware "github.com/lcortega18116/prueba/internal/middlelware"
)

func NewRouter(cfg config.Config, db *sqlx.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(chimw.RequestID)
	r.Use(chimw.RealIP)
	r.Use(middleware.ZeroLogger)
	r.Use(chimw.Recoverer)

	r.Get("/health", handlers.Health())

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/users", handlers.UsersRoutes(db))
		r.Mount("/items", handlers.ItemsRoutes(db))
	})

	log.Info().Str("env", cfg.Env).Msg("router initialized")
	return r
}
