package main

import (
	"os"
	"yml-example/config/app"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func loadConfig() {
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Configure the logger to use ConsoleWriter for pretty console output
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	log.Info().Msg("aconfig with env example starts")

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
	log.Debug().Any("primary", app.Config.Databases["primary"]).Msg("Primary Database")
	log.Debug().Any("secondary", app.Config.Databases["secondary"]).Msg("Primary Database")
}
