package routes

import (
	"tinyurl/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	r.Post("/create.tinyurl", app.TinyURLHandler.CreateTinyURL)
	r.Post("/get.tinyurl", app.TinyURLHandler.GetTinyURL)
	return r
}
