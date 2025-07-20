build:
	go build -o bin/fiber-02 server.go

mac-build:
	GOOS=darwin GOARCH=amd64 go build -o bin//fiber-02 main.go

windows-build:
	GOOS=windows GOARCH=amd64 go build -o bin//fiber-02.exe main.go

linux-build:
	GOOS=linux GOARCH=amd64 go build -o bin//fiber-02 main.go

run:
	go run cmd/main.go

test:
	go test -v -cover ./...

tidy:
	go mod tidy