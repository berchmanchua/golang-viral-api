package server

import (
	"context"
	"log"
	"net/http"

	"github.com/berchmanchua/golang-viral-api/internal/config"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func StartHTTPServer(lc fx.Lifecycle, router *chi.Mux, cfg *config.Config) {
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Starting server on", srv.Addr)
			
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("failed to listen and serve: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down HTTP server...")
			return srv.Shutdown(ctx)
		},
	})
}
