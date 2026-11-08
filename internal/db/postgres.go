package db

import (
	"context"
	"database/sql"
	
	"github.com/berchmanchua/golang-viral-api/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/fx"
)

func NewPostgresDB(lc fx.Lifecycle, cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.Postgres.DSN)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return db.PingContext(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})

	return db, nil
}
