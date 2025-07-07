package app

type DatabaseConfig struct {
	Postgres     PostgresConfig `json:"postgres"`
	Sqlite       SqliteConfig   `json:"sqlite"`
	MaxIdleConns int            `json:"maxIdleConns" default:"10" validate:"required"`
	MaxIdleTime  int            `json:"maxIdleTime" default:"30" validate:"required"` // in seconds
	MaxOpenConns int            `json:"maxOpenConns" default:"100" validate:"required"`
}

type PostgresConfig struct {
	Driver string
	Dsn    string
}

type SqliteConfig struct {
	Driver string
	Dsn    string
}
