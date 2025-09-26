package main

import (
	"context"

	"github.com/Yarik7610/effective-mobile-task/config"
	"github.com/Yarik7610/effective-mobile-task/internal/connect"
)

func main() {
	cfg := config.Load()

	connect.Postgres(context.Background(), cfg)
}
