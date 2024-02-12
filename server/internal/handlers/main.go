package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/rusmDocs/rusmDocs/internal/configs"
	"github.com/rusmDocs/rusmDocs/internal/handlers/users"
)

func InitRoute(config configs.ServerConfig) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Post("/reg", user.RegisterUser)
		r.Post("/log", user.LoginUser)
	})

	return r
}
