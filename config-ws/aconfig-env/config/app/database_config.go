package app

type DatabaseConfig struct {
	Postgres     PostgresConfig
	Sqlite       SqliteConfig
	MaxIdleConns int
	MaxIdleTime  int
	MaxOpenConns int
}

type PostgresConfig struct {
	Driver string
	Dsn    string
}

type SqliteConfig struct {
	Driver string
	Dsn    string
}
