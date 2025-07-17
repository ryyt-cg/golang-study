# sqlc-sqlite
## Getting Started
1. Create sqlc-sqlite directory
2. In the get-started directory, create a new Go module
3. Add the sqlc tool to the Go module
4. Create a new sqlc.yaml or sqlc.json file

```bash
mkdir sqlc-sqlite
cd sqlc-sqlite
go mod init sqlc-sqlite
go get github.com/sqlc-dev/sqlc/cmd/sqlc@latest
touch sqlc.yaml
```
The sqlc.yaml and sqlc.json are the configuration for sqlc generate.  It composes of:
sqlc version, sql engine, queries, package directory, out directory, etc.  For example:

```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "query.sql"
    schema: "schema.sql"
    gen:
      go:
        package: "tutorial"
        out: "tutorial"
        sql_package: "pgx/v5"
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
│   │   ├── schema.sql
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

## sclc generate
cd to pkg/db directory and run the sqlc generate command
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



