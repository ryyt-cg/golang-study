package db

import (
	"context"
	"database/sql"
)

type Database interface {
	Connect(context.Context) (*sql.DB, error)
}

type Postgres struct {
}

// Connect
// Create connection pooling using GORM postgres driver.
func (p Postgres) Connect(ctx context.Context) (*sql.DB, error) {
	//db, err := gorm.Open(postgres.Open(app.Config.Database.Postgres.Dsn), &gorm.Config{
	//	Logger: logger.Default.LogMode(logger.Info),
	//	//NamingStrategy: schema.NamingStrategy{SingularTable: true},
	//})
	//if err != nil {
	//	log.Fatal("Error connecting Postgres database.", zap.String("error", err.Error()))
	//	panic(err)
	//}
	//
	//// connection pooling configuration
	//sqlDB, _ := db.DB()
	//sqlDB.SetMaxIdleConns(app.Config.Database.MaxIdleConns)
	//sqlDB.SetMaxOpenConns(app.Config.Database.MaxOpenConns)
	//sqlDB.SetConnMaxIdleTime(time.Duration(app.Config.Database.MaxIdleTime) * time.Second)

	return nil, nil
}
