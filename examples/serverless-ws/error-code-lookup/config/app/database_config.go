package app

import validation "github.com/go-ozzo/ozzo-validation/v4"

type DatabaseConfig struct {
	Postgres     PostgresConfig
	MaxIdleConns int
	MaxOpenConns int
	MaxIdleTime  int
}

func (db DatabaseConfig) Validate() error {
	return validation.ValidateStruct(&db,
		validation.Field(&db.Postgres, validation.Required),
		validation.Field(&db.MaxIdleConns, validation.Min(0)),
		validation.Field(&db.MaxOpenConns, validation.Required),
		validation.Field(&db.MaxIdleTime, validation.Min(0)),
	)
}

type PostgresConfig struct {
	Driver   string
	Host     string
	Port     string
	Dbname   string
	Username string
	Password string
	Dsn      string
}

func (pc PostgresConfig) Validate() error {
	return validation.ValidateStruct(&pc,
		validation.Field(&pc.Driver, validation.Required),
		validation.Field(&pc.Dsn, validation.Required),
	)
}
