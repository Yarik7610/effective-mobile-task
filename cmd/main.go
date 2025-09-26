package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Yarik7610/effective-mobile-task/config"
	"github.com/Yarik7610/effective-mobile-task/internal/db/postgres"
)

func main() {
	cfg := config.Load()

	pool := postgres.Connect(context.Background(), cfg)
	postgres.RunMigrations(pool)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	<-sigs
	slog.Info("Gracefully shutdowning app")
	pool.Close()
	slog.Info("Closed Postgres pool succesfully")
}
