package app

type DatabaseConfig struct {
	Postgres     PostgresConfig
	Sqlite       SqliteConfig
	MaxIdleConns int `validate:"min=0"`
	MaxOpenConns int `validate:"min=0"`
	MaxIdleTime  int `validate:"min=0"` // in seconds
}

type PostgresConfig struct {
	Driver string `validate:"required"`
	Dsn    string `validate:"required"` // Data Source Name
}

type SqliteConfig struct {
	Driver string `validate:"required"`
	Dsn    string `validate:"required"` // Data Source Name
}
