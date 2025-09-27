package config

import "os"

type Config struct {
	ServerPort  string
	PostgresDSN string
}

func Load() *Config {
	return &Config{
		ServerPort:  getOrDefault("SERVER_PORT", "8081"),
		PostgresDSN: getOrDefault("POSTGRES_DSN", "postgres://root:password@postgres-service:5432/db?sslmode=disable"),
	}
}

func getOrDefault(envName, defaultValue string) string {
	val := os.Getenv(envName)
	if val == "" {
		val = defaultValue
	}

	return val
}
