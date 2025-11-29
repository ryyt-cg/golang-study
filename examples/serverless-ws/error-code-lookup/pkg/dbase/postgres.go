package dbase

import (
	"error-code-lookup/config/app"
	"time"

	"github.com/qiangxue/go-restful-api/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgres struct {
}

/*
PgConnect
Create connection pooling using gorm postgres driver.
*/
func PgConnect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(main.Config.Database.Postgres.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.New().Error("Error connecting Postgres database: %v", err)
		panic(err)
	}

	// connection pooling configuration
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(main.Config.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(main.Config.Database.MaxOpenConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(main.Config.Database.MaxIdleTime) * time.Second)

	return db
}
