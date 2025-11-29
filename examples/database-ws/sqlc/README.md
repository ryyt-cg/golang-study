# [SQL Compiler](https://github.com/sqlc-dev/sqlc)
SQL Compiler (sqlc) generates Go code from SQL queries. It is a tool that helps you write type-safe SQL queries in Go. It is a lightweight library that is easy to integrate into your project. It is a good choice for developers who want to write SQL queries in Go without having to write raw SQL strings.

## Supported Languages
* sqlc-gen-go
* sqlc-gen-kotlin
* sqlc-gen-python
* sqlc-gen-typescript

## How it works
sqlc generates type-safe code from SQL. Here's how it works:

1. You write queries in SQL. 
2. You run sqlc to generate code with type-safe interfaces to those queries.
3. You write application code that calls the generated code.

## Installation
Homebrew
```bash
brew install sqlc
```
go install
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

## CLI
```bash
Usage:
  sqlc [command]

Available Commands:
  compile     Statically check SQL for syntax and type errors
  completion  Generate the autocompletion script for the specified shell
  createdb    Create an ephemeral database
  diff        Compare the generated files to the existing files
  generate    Generate source code from SQL
  help        Help about any command
  init        Create an empty sqlc.yaml settings file
  push        Push the schema, queries, and configuration for this project
  verify      Verify schema, queries, and configuration for this project
  version     Print the sqlc version number
  vet         Vet examines queries

Flags:
  -f, --file string    specify an alternate config file (default: sqlc.yaml)
  -h, --help           help for sqlc
      --no-database    disable database connections (default: false)
      --no-remote      disable remote execution (default: false)

Use "sqlc [command] --help" for more information about a command.
```


