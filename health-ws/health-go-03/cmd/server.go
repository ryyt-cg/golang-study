package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/zap"
	"os"
)

var (
	r *gin.Engine
)

func loadConfig() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("health-go service starts")

	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		log.Fatal().Err(err.Error()).
			Msg("Fail to load application configuration.")
	}

	//pg = dbase.PgConnect()
	var err error
	sqlite, err = dbase.SqliteConnect()
	if err != nil {
		logger.Fatal("Fail to connect the database.",
			zap.String("error", err.Error()))
		//os.Exit(-1)
	}

	// Creates a router without any middleware by default
	r = gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default, gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
}

// Component initialization
func loadComponents() {
	log.Debug().Msg("Component initialization starts")

	home := r.Group("/")
	v1 := r.Group("/v1")
	//v1.Use(middleware.Authenticate(authenService))

	healthCheckRouter.Register(home.Group("/health"))
}

//	@title			Pet Clinic API
//	@version		1.0
//	@description	This is a pet clinic API server.

//		@Schemes	http
//	 @Host		localhost:8080
func main() {
	loadConfig()
	loadComponents()
	// add swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//r.Run(":8080")
	httpServer := ds.NewHttpServer(r)
	httpRouter := httpServer.HttpRouter()

	g.Go(func() error {
		return httpRouter.ListenAndServeTLS(app.Config.Server.CertFile, app.Config.Server.KeyFile)
	})

	if err := g.Wait(); err != nil {
		logger.Fatal("Fail to run http server.", zap.String("error", err.Error()))
		os.Exit(-1)
	}
}
