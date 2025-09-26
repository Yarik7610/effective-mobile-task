package postgres

import (
	"context"
	"log"
	"log/slog"

	"github.com/Yarik7610/effective-mobile-task/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("Can't connect to Postgres: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Can't Ping Postgres: %v", err)
	}

	slog.Info("Connected to Postgres succesfully")
	return pool
}
