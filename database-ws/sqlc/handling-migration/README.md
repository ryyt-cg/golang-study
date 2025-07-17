# Handling SQL migrations
sqlc does not perform database migrations for you. However, sqlc is able to differentiate between up and down migrations. sqlc ignores down migrations when parsing SQL files.

sqlc supports parsing migrations from the following tools:
* atlas
* dbmate
* golang-migrate
* goose
* sql-migrate
* tern

```bash
mkdir handling-migration
cd handling-migration
go mod init handling-migration
go get github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```
The sqlc.yaml and sqlc.json are the configuration for sqlc generate.  It composes of:
sqlc version, sql engine, queries, package directory, out directory, etc.  For example:


In this example, sqlc refers schema files from the `migrations` directory, which contains SQL migration files. The queries are defined in the `query` directory, and the generated code will be placed in the `sqlgen` directory.

```yaml
version: "2"
sql:
  - engine: "sqlite"
    queries:
      - "query/author.sql"
      - "query/book.sql"
    schema: "../../migrations"
    gen:
      go:
        package: "sqlgen"
        out: "sqlgen"
```

## Project Structure
```bash
├── Makefile
├── README.md
├── cmd
│   └── server.go
├── go.mod
├── go.sum
├── migrations
│   └── 00001_initial_schema.sql
├── pkg
│   ├── db
│   │   ├── dbase.go
│   │   ├── postgres.go
│   │   ├── query
│   │   │   ├── author.sql
│   │   │   └── book.sql
│   │   ├── sqlc.yaml
│   │   ├── sqlgen
│   │   │   ├── author.sql.go
│   │   │   ├── book.sql.go
│   │   │   ├── db.go
│   │   │   └── models.go
│   │   └── sqlite.go
│   └── repository
│       ├── author_repository.go
│       └── book_repository.go
└── sqlc-tutorial.sqlite
```

## Create sqlite database
Use goose (database migration tool) to create the sqlc-tutorial.sqlite database.
```bash
go get -u github.com/pressly/goose/v3/cmd/goose
make goose-up
```

## sqlc generate
cd to pkg/db directory and run the `sqlc generate` command
```bash
cd pkg/db
sqlc generate
```
This will generate sqlgen directory with the sqlc generated code.
- author.sql.go contains the generated code for author queries.
- book.sql.go contains the generated code for book queries.
- db.go contains the generated code for the database connection and transaction management.
- models.go contains the generated code for the database models.

## Run the application
```bash
go run cmd/server.go
```



