package main

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestGetEcodeById(t *testing.T) {
	ecodeMock := main.MockEcodeServicer{}
	ecodeRespA := &main.EcodeResponse{
		ID:          "DIP-TEST",
		Description: "testing connection",
	}

	ecodeMock.On("GetEcodeById", "DIP-TEST").Return(ecodeRespA, nil)
	logger := log.New().With(context.TODO(), "version", "test")
	ecodeFunction := main.NewEcodeFunction(logger, &ecodeMock)

	t.Run("Successful Request", func(t *testing.T) {
		request := events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"id": "DIP-TEST",
			},
			Resource:   "/v1/ecode/{id}",
			HTTPMethod: "GET",
		}
		response, err := ecodeFunction.handler(context.TODO(), request)
		if err != nil {
			t.Fatal("Everything should be ok")
		}

		assert.Equal(t, 200, response.StatusCode)
		var ecodeRespE main.EcodeResponse
		_ = json.Unmarshal([]byte(response.Body), &ecodeRespE)
		assert.Equal(t, ecodeRespE.ID, ecodeRespA.ID)
		assert.Equal(t, ecodeRespE.Description, ecodeRespA.Description)
	})
}

func TestEcodeNotFound(t *testing.T) {
	ecodeMock := main.MockEcodeServicer{}

	ecodeMock.On("GetEcodeById", "DIP-FAIL").Return(nil, nil)
	logger := log.New().With(context.TODO(), "version", "test")
	ecodeFunction := main.NewEcodeFunction(logger, &ecodeMock)

	t.Run("Not Found Request", func(t *testing.T) {
		request := events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"id": "DIP-FAIL",
			},
			Resource:   "/v1/ecode/{id}",
			HTTPMethod: "GET",
		}
		response, err := ecodeFunction.handler(context.TODO(), request)
		if err != nil {
			t.Fatal("Everything should be ok")
		}

		assert.Equal(t, 404, response.StatusCode)
		assert.Equal(t, "{}", response.Body)
	})
}

func TestEcodeHandlerFail(t *testing.T) {
	ecodeMock := main.MockEcodeServicer{}

	ecodeMock.On("GetEcodeById", "DIP-FAIL").Return(nil, errors.New("connection timeout"))
	logger := log.New().With(context.TODO(), "version", "test")
	ecodeFunction := main.NewEcodeFunction(logger, &ecodeMock)

	t.Run("Failed Request", func(t *testing.T) {
		request := events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"id": "DIP-FAIL",
			},
			Resource:   "/v1/ecode/{id}",
			HTTPMethod: "GET",
		}
		response, err := ecodeFunction.handler(context.TODO(), request)
		if err != nil {
			t.Fatal("Everything should be ok")
		}

		assert.Equal(t, 500, response.StatusCode)
		assert.Equal(t, "{}", response.Body)
	})
}
