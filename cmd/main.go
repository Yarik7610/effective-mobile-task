package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Yarik7610/effective-mobile-task/config"
	"github.com/Yarik7610/effective-mobile-task/db/postgres"
	"github.com/Yarik7610/effective-mobile-task/internal/transport/http"
)

func main() {
	cfg := config.Load()

	pool := postgres.Connect(context.Background(), cfg)
	postgres.RunMigrations(pool)

	router := http.InitRouter(pool)
	if err := router.Run(fmt.Sprintf(":%s", cfg.ServerPort)); err != nil {
		log.Fatalf("Can't run router: %v", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	<-sigs
	slog.Info("Gracefully shutdowning app")
	pool.Close()
	slog.Info("Closed Postgres pool succesfully")
}
