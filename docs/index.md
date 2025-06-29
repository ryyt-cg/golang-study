# Welcome to My Golang Study Note

This repository is a collection of my notes and examples while learning Go & its support libraries. It covers various technical stacks, such as configuration, logging, messaging, ORM, route, etc.  I create one workspace each tech stack, each tech stack covers several different library implementation modules.  For example, logging workspace contains three different libraries, such as uber/zap, zerolog, and slog. The examples are designed to be simple and easy to understand, focusing on the core concepts of each topic. The code is written in a modular way, allowing you to run each example independently.

## Restful API Project/Module Layout

```bash
├── Dockerfile
├── Makefile
├── README.md
├── api                                                 // API layer
│   ├── health                                          // Health check API
│   │   ├── health.go
├── `command`                                           // Command line interface
│   ├── certs                                           // Certificate generation for HTTPS
│   │   ├── myCA-without-passphrase.key
│   │   ├── myCA.key
│   │   └── myCA.pem
│   └── server.go                                       // Main entry point for the application
├── config                                              // Configuration files
│   ├── app                                             // Application configuration
│   │   ├── config.go
│   │   ├── database_config.go
│   │   ├── info_config.go
│   │   ├── server_config.go
│   ├── app-dev.yaml
│   ├── app-prd.yaml                                    // Production configuration file
│   ├── app-stg.yaml
│   ├── app.yaml                                        // Default configuration file
│   └── errors.yaml
├── doc
├── docs                                               // Swagger documentation
│   ├── docs.go
│   ├── petclinic-ermodel.png
│   ├── petclinic-model.png
│   ├── swagger.json
│   ├── swagger.yaml
│   └── test-driven-development-golang.png
├── go.mod
├── go.sum
├── internal
├── middleware                                         // Middleware layer
│   ├── errors
│   │   ├── middleware.go
│   │   ├── response.go
│   │   └── response_test.go
│   ├── request_header.go
│   └── request_header_test.go
├── migrations                                        // Database migrations
│   ├── postgres
│   │   └── 2020092012000_petclinic.sql
│   └── sqlite
│       └── 2020092012000_petclinic.sql
├── pkg                                               // Shared packages
│   ├── accesslog
│   │   ├── middleware.go
│   │   └── middleware_test.go
│   ├── dbase
│   │   ├── postgres.go
│   │   ├── sqlite.go
│   ├── ds
│   │   ├── http_server.go
│   │   └── http_server_test.go
│   ├── infra
│   │   └── repository
│   │       ├── owner.go
│   │       ├── owner_repository.go
│   │       ├── pet.go
│   │       ├── pet_repository.go
│   │       ├── vet.go
│   │       ├── vet_repository.go
│   │       ├── visit.go
│   ├── log
│   │   ├── logger.go
│   │   ├── logger_test.go
│   │   └── mock_logger.go
│   ├── model
│   │   ├── context.go
│   │   ├── person.go
│   │   └── person_test.go
```

