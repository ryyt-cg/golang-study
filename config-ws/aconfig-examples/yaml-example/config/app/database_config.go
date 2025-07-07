package app

type DatabaseConfig struct {
	Postgres     PostgresConfig `yaml:"postgres"`
	Sqlite       SqliteConfig   `yaml:"sqlite"`
	MaxIdleConns int            `yaml:"maxIdleConns" default:"10" validate:"required"`
	MaxIdleTime  int            `yaml:"maxIdleTime" default:"30" validate:"required"` // in seconds
	MaxOpenConns int            `yaml:"maxOpenConns" default:"100" validate:"required"`
}

type PostgresConfig struct {
	Driver string
	Dsn    string
}

type SqliteConfig struct {
	Driver string
	Dsn    string
}
