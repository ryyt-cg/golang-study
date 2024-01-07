package main

import (
	"gitlab.con/aionx/go-examples/grpc-in-go/order-service/config/app"
	"gitlab.con/aionx/go-examples/grpc-in-go/order-service/pkg/db"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

var (
	log *zap.Logger
	pg  *gorm.DB
)

// Instantiate zap log
// Instantiate gin router and middlewares
func init() {

	//zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	log, _ = zap.NewProduction()
	log.Info("gRCP in Go starts")

	// load application configurations
	if err := app.LoadConfig("./order-service/config"); err != nil {
		log.Error("failed to load application configuration.",
			zap.String("error", err.Error()))
		os.Exit(-1)
	}

	var exception error
	pg = db.PgConnect()
	if exception != nil {
		log.Fatal("Order Database fails to connect.")
	}
}

func main() {

}
