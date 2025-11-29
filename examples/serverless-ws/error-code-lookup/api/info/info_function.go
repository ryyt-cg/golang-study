package info

import (
	"context"
	"encoding/json"
	main2 "error-code-lookup/config/app"
	errors2 "error-code-lookup/pkg/errors"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	ddlambda "github.com/DataDog/datadog-lambda-go"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/qiangxue/go-restful-api/pkg/log"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

type AppInfoServicer interface {
	getAppInfo() (*info.AppInfo, error)
}

type AppInfoFunction struct {
	infoService AppInfoServicer
}

func NewAppInfoFunction(infoService AppInfoServicer) *AppInfoFunction {
	return &AppInfoFunction{infoService: infoService}
}

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("no IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("non 200 Response found")
)

/*
 * Last function to call the lambda logs & traces
 */
//func (function *AppInfoFunction) ddTracing(ctx context.Context) {
//	currentSpan, _ := tracer.SpanFromContext(ctx)
//	function.logger.Infof("info log message: %v", currentSpan)
//}

/*
 * handler must return both response and error.
 */
func (function *AppInfoFunction) handler(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	// Trace an HTTP request
	//req, _ := http.NewRequestWithContext(ctx, "GET", "https://www.datadoghq.com", nil)
	client := http.Client{}
	client = *httptrace.WrapClient(&client)
	//client.Do(req)

	// Connect your Lambda logs and traces
	//defer function.ddTracing(ctx)

	resp, err := http.Get(DefaultHTTPGetAddress)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if resp.StatusCode != 200 {
		return events.APIGatewayProxyResponse{}, ErrNon200Response
	}

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if len(ip) == 0 {
		return events.APIGatewayProxyResponse{}, ErrNoIP
	}

	appInfo, err := InfoService.getAppInfo()
	var body []byte

	if err != nil {
		body, _ := json.Marshal(errors2.InternalServerError(err.Error()))

		return events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 500,
		}, err
	} else {
		appInfo = string(ip)
		body, _ = json.Marshal(appInfo)

		return events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 200,
		}, nil
	}
}

func main() {
	// load application configurations
	if err := main2.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	logger := log.New().With(context.TODO(), "version", main2.Config.AppInfo.Version)
	logger.Infof("App Info function loading...")

	appInfoService := NewService()
	appInfoFunction := NewAppInfoFunction(appInfoService)

	// Wrap your handler function
	lambda.Start(ddlambda.WrapHandler(appInfoFunction.handler, nil))
	//lambda.Start(ddlambda.WrapHandler(appInfoFunction.handler, &ddlambda.Config{
	//	BatchInterval: time.Second * 15,
	//	APIKey:        "api-key",
	//}))
}
