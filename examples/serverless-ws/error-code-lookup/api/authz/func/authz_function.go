package main

import (
	"context"
	"encoding/json"
	"error-code-lookup/api/authz"
	"error-code-lookup/config/app"
	"error-code-lookup/pkg/errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/qiangxue/go-restful-api/pkg/log"
)

type AuthenServicer interface {
	Verify(bearerToken string) error
}

type AuthzFunction struct {
	authService AuthenServicer
}

func NewAuthzFunction(authService AuthenServicer) *AuthzFunction {
	return &AuthzFunction{authService: authService}
}

func (function *AuthzFunction) handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	authHeader := request.Headers["Authorization"]
	err := function.authService.Verify(authHeader)
	if err != nil {
		body, _ := json.Marshal(errors.Unauthorized(err.Error()))
		return events.APIGatewayProxyResponse{
			StatusCode: 401,
			Body:       string(body),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "good token",
	}, nil
}

func main() {
	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	logger := log.New().With(context.TODO(), "version", app.Config.AppInfo.Version)
	logger.Info("Verifier function loading...")

	authenService := authz.NewAuthenService()
	verifierFunction := NewAuthzFunction(authenService)

	lambda.Start(verifierFunction.handler)
}
