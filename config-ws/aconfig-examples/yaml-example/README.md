# Application Configuration Using YAML and TOML Example
Usage of `aconfig` library in Go to manage application configuration.

Precedence or order of configuration sources:
1. defaults in the code
2. files (JSON, YAML, TOML, DotENV, HCL)
3. environment variables
4. command-line flags

In  this example, these sources are used:
* default values
* YAML config file
* environment variables


Project structure:
```
├── README.md
├── cmd
│   └── server.go
├── config
│   ├── app
│   │   ├── config.go
│   │   ├── database_config.go
│   │   ├── info_config.go
│   │   └── server_config.go
│   ├── app-dev.yaml
│   ├── app-prd.yaml
│   ├── app.yaml
│   └── example_config.yaml
├── go.mod
├── go.sum
└── simple-toml-example.go  // Example of using TOML configuration
```


Database configuration example:
* Use `yaml:"attributeName"` to specify the YAML attribute name for each field.
* Use `default:"value"` to set a default value for the field.
* Use `validate:"required"` to ensure the field is required and must be provided.
* If `yaml:"attributeName"` is not specified, the underscore name will be applied. i.e. `MaxIdleConns` will be serialized as `max_idle_conns` in YAML.

Note: Need to install YAML and/or toml decoder library:
```bash
go get github.com/cristalhq/aconfig/aconfigtoml
go get github.com/cristalhq/aconfig/aconfigtoml
````

```go
type DatabaseConfig struct {
	Postgres     PostgresConfig `yaml:"postgres"`
	Sqlite       SqliteConfig   `yaml:"sqlite"`
	MaxIdleConns int            `default:"10" validate:"required"`
	MaxOpenConns int            `default:"100" validate:"required"`
	MaxIdleTime  int            `yaml:"maxIdleTime" default:"30" validate:"required"` // in seconds
}
```

YAML configuration file will look like this:
MaxIdleTime is not specified, so it will use the default value of 30 seconds.

```yaml
database:
  sqlite:
    driver: sqlite3
    dsn: ${SQLITE_DSN}
  postgres:
    driver: postgres
    dsn: ${POSTGRES_DSN}
  max_idle_conns: 1
  max_open_Conns: 5
```

* Use environment variables to override the configuration values:
  The environment variable nomenclature is `DATABASE_POSTGRES_DSN` for Postgres and `DATABASE_SQLITE_DSN` for SQLite, because they reside in the nest structs of `DatabaseConfig` and PostgresConfig and SqliteConfig respectively.  If attribute names are not correctly set, then values will not override the defaults or yaml file values.

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


**Troubleshooting:**<br/>
If you receive error logs like below, it means that the configuration file is not loaded correctly or has unknown fields.
```bash
0:56AM INF aconfig with env example starts
./config/app.yaml
10:56AM FTL Fail to load application configuration. error="failed to load configuration file ./config/app.yaml: load config: load files: unknown field in file \"./config/app.yaml\": database.maxIdleConns (see AllowUnknownFields config param)"
```