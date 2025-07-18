package main

import (
	"fiber-01/api/author"
	"fiber-01/api/info"
	"fiber-01/config/app"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	fiberApp *fiber.App
)

// Instantiate zerolog
// Instantiate fiber router and middlewares
func loadConfig() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Go Fiber 01 starts")

	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		log.Fatal().Err(err).
			Msg("Fail to load application configuration.")
	}

	// Create a new Fiber instance
	fiberApp = fiber.New()
	// Apply global middlewares
	fiberApp.Use(healthcheck.New())
	fiberApp.Use(recover.New())   // Recover from panics and continue
	fiberApp.Use(requestid.New()) // Generate a unique request ID for each request
}

func loadComponents() {
	infoService := info.NewService()
	infoRouter := info.NewRouter(infoService)

	authorService := author.NewService()
	authorRouter := author.NewRouter(authorService)

	// create a new group for the /api/gof endpoint
	home := fiberApp.Group(app.Config.Server.BaseURL)
	// Register the info router to the home group
	infoRouter.Register(home.Group("/info"))

	// create a new group for the /api/gof/v1 endpoint
	v1 := fiberApp.Group(app.Config.Server.BaseURL + "/v1")
	// Register the author router to the v1 group
	authorRouter.Register(v1.Group("/authors"))
}

func main() {
	loadConfig()
	loadComponents()

	// Start the server on port 3000
	err := fiberApp.Listen(app.Config.Server.HttpPort)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
		return
	}
}
