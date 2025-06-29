package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"viper/config/app"
)

func loadConfig() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("viper example starts")

	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		log.Fatal().Err(err).
			Msg("Fail to load application configuration.")
	}
}

func main() {
	loadConfig()
	log.Debug().Any("server", app.Config.Server).Msg("Server Config")
	log.Debug().Any("appInfo", app.Config.AppInfo).Msg("App Version")
	log.Debug().Any("sqlite", app.Config.Database.Sqlite).Msg("SQLite Config")
	log.Debug().Any("postgres", app.Config.Database.Postgres).Msg("Postgres Config")
}
