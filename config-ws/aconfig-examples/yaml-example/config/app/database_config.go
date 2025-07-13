package app

type DatabaseConfig struct {
	Postgres     PostgresConfig `yaml:"postgres"`
	Sqlite       SqliteConfig   `yaml:"sqlite"`
	MaxIdleConns int            `default:"10" validate:"required"`
	MaxOpenConns int            `default:"100" validate:"required"`
	MaxIdleTime  int            `yaml:"maxIdleTime" default:"30" validate:"required"` // in seconds
}

type PostgresConfig struct {
	Driver string
	Dsn    string
}

type SqliteConfig struct {
	Driver string
	Dsn    string
}
