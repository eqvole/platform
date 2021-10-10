package rest

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/eqvole/platform/pkg/rest/handlers"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(30 * time.Second))

	userhandler := handlers.NewUserHandler()
	r.Route("/user", func(r chi.Router) {
		r.Post("/register", userhandler.Register)
		r.Post("/auth", userhandler.Auth)
	})
	return r
}
