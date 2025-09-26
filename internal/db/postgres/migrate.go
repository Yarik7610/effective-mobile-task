package postgres

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

const POSTGRES_MIGRATIONS_PATH = "internal/db/postgres/migrations"

func RunMigrations(pool *pgxpool.Pool) {
	db := stdlib.OpenDBFromPool(pool)
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Can't prepare Driver for Postgres Mirator: %v", err)
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", POSTGRES_MIGRATIONS_PATH),
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("Can't create Migrator for Postgres: %v", err)
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	slog.Info("Migrations to Postgres applied successfully")
}
