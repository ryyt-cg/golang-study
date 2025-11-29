.PHONY: build

build:
	go-bindata -o config/app/app.yaml.go config/...
	sam build

deploy-sbx: build
	sam deploy --profile aws-digital-sbx-Broad-access-role  --stack-name error-code-lookup  --config-env DEV-WEST --region us-west-1  --confirm-changeset

deploy: build
	sam deploy --profile digital-oz-DigitalDeveloper-CXT-DIP  --stack-name error-code-lookup  --config-env DEV-WEST --region us-west-1  --confirm-changeset

local: build
	sam local start-api

local-debug: build
	sam local start-api -d 5986 --debugger-path ~/go/debug --debug-args "-delveAPI=2"

mock_ecode:
	mockery --case=underscore --dir api/ecode --name=EcodeRepositorier  --inpackage
	mockery --case=underscore --dir api/ecode --name=EcodeServicer  --inpackage

debug_authz:
	GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o api/authz/function/authz api/authz/function/authz_function.go

debug_info:
	GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o api/info/function/info api/info/function/info_function.go

debug_ecode:
	GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o api/ecode/function/ecode api/ecode/func/ecode_function.go

build-AuthzFunction:
	GOOS=linux GOARCH=arm64 go build -o authz_function ./api/authz/func/authz_function.go
	cp ./authz_function $(ARTIFACTS_DIR)/bootstrap

build-AppInfoFunction:
	GOOS=linux GOARCH=arm64 go build -o info_function ./api/info/func/info_function.go
	cp ./info_function $(ARTIFACTS_DIR)/bootstrap

build-InsertEcodeFunction:
	GOOS=linux GOARCH=arm64 go build -o ecode_function ./api/info/func/ecode_function.go
	cp ./ecode_function $(ARTIFACTS_DIR)/bootstrap

build-GetAllEcodesFunction:
	GOOS=linux GOARCH=arm64 go build -o ecode_function ./api/info/func/ecode_function.go
	cp ./ecode_function $(ARTIFACTS_DIR)/bootstrap


#sam local start-api -d 5986 --debugger-path  /Users/ryyt/go/debug --debug-args "-delveAPI=2"
#sam local invoke SqsSubFunction --debug-port 5986 --debugger-path /Users/ryyt/go/debug --debug-args "-delveAPI=2" --event events/sqs-receive-message-request.json
