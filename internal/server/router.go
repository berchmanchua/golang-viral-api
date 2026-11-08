package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(viralHandler *ViralHandler) *chi.Mux {
	r := chi.NewRouter()
	
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1/viral", func(r chi.Router) {
		r.Post("/", viralHandler.Create)
		r.Get("/{id}", viralHandler.Get)
		r.Put("/{id}/views", viralHandler.UpdateViews)
		r.Delete("/{id}", viralHandler.Delete)
	})
	
	return r
}
