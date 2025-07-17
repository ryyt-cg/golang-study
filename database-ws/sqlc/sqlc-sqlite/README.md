


## Getting Started
1. Create 01-getting-started directory
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


