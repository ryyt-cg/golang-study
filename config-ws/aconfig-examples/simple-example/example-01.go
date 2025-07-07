package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"simple-example/config/app"
)

func loadConfig() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("aconfig with env example starts")

	// load application configurations
	if err := app.LoadConfig("config"); err != nil {
		log.Fatal().Err(err).
			Msg("Fail to load application configuration.")
	}
}

func main() {
	// Set environment variables for testing
	// In a real application, these would be set in the environment or through a configuration management system.
	os.Setenv("DATABASE_POSTGRES_DSN", "env-postgres-dsn")
	os.Setenv("DATABASE_SQLITE_DSN", "env-sqlite-dsn")
	defer os.Clearenv()

	loadConfig()
	log.Info().Any("appInfo", app.Config.AppInfo).Msg("App Info")
	log.Info().Any("server", app.Config.Server).Msg("Server Info")
	log.Info().Any("database", app.Config.Database).Msg("Database Info")
}
