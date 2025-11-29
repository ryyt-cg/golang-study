# [Application Configuration](https://github.com/avelino/awesome-go?tab=readme-ov-file#configuration)
Application configuration is an essential part of any software project. In Go, configuration files can be managed using various methods, such as environment variables, JSON, YAML, or TOML files. The configuration files typically include settings for the application, database connections, server settings, and other environment-specific parameters.

## [aconfig](https://github.com/cristalhq/aconfig)
Aconfig is a simple and powerful configuration library for Go. It supports various formats like JSON, YAML, TOML, and more. Aconfig allows you to define your configuration structure using Go structs and automatically maps the configuration values to the struct fields.

### Features:
* Automatic fields mapping.
* Supports different sources:
    * defaults in the code
    * files (JSON, YAML, TOML, DotENV, HCL)
    * environment variables
    * command-line flags
* dependency-free (file parsers are optional).

Two examples of using Aconfig are provided below:
* [Aconfig Examples](aconfig-examples/README.md)
* [YAML File Example](config-ws/aconfig-examples/yaml-example/README.md)

## [godotenv](https://github.com/joho/godotenv)
dotenv is a Go library that allows you to load environment variables from a `.env` file into your Go application. This is particularly useful for managing configuration settings in a development environment without hardcoding them into your source code.


## [Viper](https://github.com/spf13/viper)
Viper is a popular and complete configuration solution for Go applications. It supports reading from JSON, TOML, YAML, HCL, and Java properties files. Viper can also read from environment variables and remote configuration systems like Consul and etcd. It provides a flexible way to manage application configurations and allows you to set default values, read from multiple sources, and watch for changes in configuration files.

