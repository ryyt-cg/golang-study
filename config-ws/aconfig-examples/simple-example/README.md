# Application Configuration Using Built-in JSON
Usage of `aconfig` library in Go to manage application configuration.

Precedence or order of configuration sources:
1. defaults in the code
2. files (JSON, YAML, TOML, DotENV, HCL)
3. environment variables
4. command-line flags

In  this example, these sources are used:
* default values
* JSON config file (built-in `aconfig` support)
* environment variables

Project structure:
```plaintext
├── README.md
├── config
│   ├── app
│   │   ├── config.go
│   │   ├── database_config.go
│   │   ├── info_config.go
│   │   └── server_config.go
│   └── app.json
├── example-01.go
├── go.mod
└── go.sum
```


Database configuration example:
* Use `json:"attributeName"` to specify the JSON attribute name for each field.
* Use `default:"value"` to set a default value for the field.
* Use `validate:"required"` to ensure the field is required and must be provided.
* If `json:"attributeName"` is not specified, the underscore name will be applied. i.e. `MaxIdleConns` will be serialized as `max_idle_conns` in JSON.

```go
type DatabaseConfig struct {
	Postgres     PostgresConfig `json:"postgres"`
	Sqlite       SqliteConfig   `json:"sqlite"`
	MaxIdleConns int            `default:"10" validate:"required"`
	MaxOpenConns int            `default:"100" validate:"required"`
	MaxIdleTime  int            `json:"maxIdleTime" default:"30" validate:"required"` // in seconds
}
```

JSON configuration file will look like this:
- MaxIdleTime is not specified, so it will use the default value of 30 seconds.

```json
{
  "postgres": {
    "driver": "postgres",
    "dsn": "${POSTGRES_DSN}"
  },
  "sqlite": {
    "driver": "sqlite3",
    "dsn": "${SQLITE_DSN}"
  },
  "max_idle_conns": 1,
  "max_open_conns": 5
}
```

* Use environment variables to override the configuration values:
The environment variable nomenclature is `DATABASE_POSTGRES_DSN` for Postgres and `DATABASE_SQLITE_DSN` for SQLite, because they reside in the nest structs of `DatabaseConfig` and PostgresConfig and SqliteConfig respectively.  If attribute names are not correctly set, then values will not override the defaults or JSON file values.

Three ways to set environment variables:
1. By terminal (Linux, macOS). 
```bash
export DATABASE_POSTGRES_DSN="postgres://user:password@localhost:5432/dbname?sslmode=disable"
export DATABASE_POSTGRES_DSN="file:mydb.sqlite?cache=shared&mode=rwc"
```
2. In the go code:
```go
	os.Setenv("DATABASE_POSTGRES_DSN", "env-postgres-dsn")
	os.Setenv("DATABASE_POSTGRES_DSN", "env-sqlite-dsn")
```
3. In IDE such as Goland, Intellij, Visual Studio Code, etc. 
   - Go to Run/Debug Configurations
   - Add environment variables in the Environment Variables field.

