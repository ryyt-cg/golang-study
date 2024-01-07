package db

import (
	"gitlab.con/aionx/go-examples/grpc-in-go/order-service/config/app"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type Postgres struct {
}

// PgConnect
// Create connection pooling using gorm postgres driver.
func PgConnect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(app.Config.Database.Postgres.Dsn), &gorm.Config{
		//Logger: zap.L(),
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		zap.L().Error("Error connecting Postgres database", zap.String("error", err.Error()))
		panic(err)
	}

	// connection pooling configuration
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(app.Config.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(app.Config.Database.MaxOpenConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(app.Config.Database.MaxIdleTime) * time.Second)

	return db
}
