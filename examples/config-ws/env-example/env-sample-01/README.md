# [GoDotEnv](https://github.com/joho/godotenv)
Dotenv is a Go library that allows you to load environment variables from a `.env` file into your Go application. This is particularly useful for managing configuration settings in a development environment without hardcoding them into your source code.

## Install GoDotEnv
```bash
go get github.com/joho/godotenv
```

## Usage
1. create a `.env` file in your project root directory with key-value pairs:
   - KEY_NAME=value
   - ANOTHER_KEY=another_value
2. import the package in your Go code:
```go
   "github.com/joho/godotenv"
````
3. load the environment variables from the `.env` file:
```go
   err := godotenv.Load()
   if err != nil {
       log.Fatal("Error loading .env file")
   }
```
4. access the environment variables using `os.Getenv`:
```go
   value := os.Getenv("KEY_NAME")
   fmt.Println(value)
```

Another way to take advantage of the autoload package.
1. Import the autoload package:
```go
   "github.com/joho/godotenv/autoload"
```
2. access the environment variables as usual:`
```go
   value := os.Getenv("KEY_NAME")
   fmt.Println(value)
```

