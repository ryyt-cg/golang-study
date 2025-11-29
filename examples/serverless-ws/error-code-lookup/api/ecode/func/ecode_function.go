package main

import (
	"context"
	"encoding/json"
	"error-code-lookup/config/app"
	"error-code-lookup/pkg/dbase"
	"error-code-lookup/pkg/errors"
	"fmt"
	"net/http"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"error-code-lookup/api/ecode"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/qiangxue/go-restful-api/pkg/log"
)

type EcodeFunction struct {
	ecodeService ecode.EcodeServicer
}

func NewEcodeFunction(ecodeService ecode.EcodeServicer) *EcodeFunction {
	return &EcodeFunction{ecodeService: ecodeService}
}

func (function *EcodeFunction) handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Trace an HTTP request
	//req, _ := http.NewRequestWithContext(ctx, "GET", "https://www.datadoghq.com", strings.NewReader(request.Body))
	client := http.Client{}
	client = *httptrace.WrapClient(&client)
	//client.Do(req)

	var response events.APIGatewayProxyResponse
	log.Info().Str("URL", request.Resource).Msg("Received request")

	if request.HTTPMethod == http.MethodGet {
		if request.Resource == "/v1/ecode/{id}" {
			response = function.ecodeById(request)
		} else if request.Resource == "/v1/ecode/all" {
			response = function.ecodeAll()
		}
	} else if request.HTTPMethod == http.MethodPut {
		response = function.updateEcode(request)
	} else if request.HTTPMethod == http.MethodPost {
		response = function.insertEcode(request)
	} else if request.HTTPMethod == http.MethodDelete {
		response = function.deleteEcode(request)
	}

	// Connect your Lambda logs and traces
	currentSpan, _ := tracer.SpanFromContext(ctx)
	log.Infof("info log message: %v", currentSpan)
	return response, nil
}

func (function *EcodeFunction) ecodeAll() events.APIGatewayProxyResponse {
	log.Infof("retrieve all ecodes")
	result, err := function.ecodeService.GetAllEcodes()
	if err != nil {
		log.Errorf("Unable to retrieve all ecodes:", err.Error())

		body, _ := json.Marshal(errors.InternalServerError(err.Error()))
		return events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 500,
		}
	}

	body, _ := json.Marshal(result)
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}
}

func (function *EcodeFunction) ecodeById(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	code := request.PathParameters["id"]
	log.Infof("ecode function path param: %v", code)

	result, err := function.ecodeService.GetEcodeById(code)
	var body []byte

	if result == nil {
		body, _ := json.Marshal(errors.NotFound(err.Error()))
		return events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 404,
		}
	}

	if err != nil {
		log.Errorf("GetEcodeById('%v') fails", code)
		body, _ = json.Marshal(errors.InternalServerError(""))

		return events.APIGatewayProxyResponse{
			Body:       string("body"),
			StatusCode: 500,
		}
	}

	body, _ = json.Marshal(result)

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}
}

func (function *EcodeFunction) updateEcode(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	id := request.PathParameters["id"]
	log.Infof("ecode function path param: %v", id)

	var ecodeReq ecode.Ecode
	err := json.Unmarshal([]byte(request.Body), &ecodeReq)
	if err != nil {
		log.Errorf("unmarshal fails: %v", err)

		body, _ := json.Marshal(errors.BadRequest(err.Error()))
		return events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 400,
		}
	}

	if ecodeReq.ID != id {
		log.Errorf("request code: %v not match path code: %v", ecodeReq.ID, id)

		body, _ := json.Marshal(errors.BadRequest(err.Error()))
		return events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 400,
		}
	}

	result, err := function.ecodeService.Update(&ecodeReq)
	if err != nil {
		log.Errorf("Unable to update ecode: %s", err.Error())

		body, _ := json.Marshal(errors.InternalServerError(err.Error()))
		return events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 500,
		}
	}

	body, _ := json.Marshal(result)
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}
}

func (function *EcodeFunction) insertEcode(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var ecodeReq ecode.Ecode
	err := json.Unmarshal([]byte(request.Body), &ecodeReq)
	if err != nil {
		log.Errorf("unmarshal fails: %v", err)

		body, _ := json.Marshal(errors.BadRequest(err.Error()))
		return events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 400,
		}
	}

	result, err := function.ecodeService.Insert(&ecodeReq)
	if err != nil {
		log.Errorf("Unable to insert ecode: %s", err.Error())

		body, _ := json.Marshal(errors.InternalServerError(err.Error()))
		return events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 500,
		}
	}

	body, _ := json.Marshal(result)
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}
}

func (function *EcodeFunction) deleteEcode(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	codeId := request.PathParameters["id"]
	log.Infof("ecode function path param: %v", codeId)

	_, err := function.ecodeService.DeleteById(codeId)
	if err != nil {
	}

	return events.APIGatewayProxyResponse{
		Body:       "{}",
		StatusCode: 200,
	}
}

func main() {
	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	logger := log.New().With(context.TODO(), "version", app.Config.AppInfo.Version)
	logger.Info("Ecode function loading...")
	pg := dbase.PgConnect()

	ecodeRepository := ecode.NewEcodeRepository(pg)
	ecodeService := ecode.NewEcodeService(ecodeRepository)
	ecodeFunction := NewEcodeFunction(ecodeService)

	lambda.Start(ecodeFunction.handler, nil)
}
