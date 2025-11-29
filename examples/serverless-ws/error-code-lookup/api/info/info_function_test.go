package info

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/qiangxue/go-restful-api/pkg/log"
	"github.com/stretchr/testify/mock"
)

type appInfoServiceMock struct {
	mock.Mock
}

func (ism *appInfoServiceMock) GetAppInfo() (*main.main.AppInfo, error) {
	args := ism.Called()
	intf := args.Get(0)
	val := intf.(main.main.AppInfo)
	return &val, args.Error(1)
}

func TestInfoHandler(t *testing.T) {
	infoMock := appInfoServiceMock{}
	infoA := &main.main.AppInfo{
		AppName:     "Test App",
		Description: "Info App",
		Version:     "1.0.0",
		Ip:          "127.0.0.1",
	}
	infoMock.On("GetAppInfo").Return(*infoA, nil)

	logger := log.New().With(context.TODO(), "version", "test")
	infoFunction := NewAppInfoFunction(logger, &infoMock)

	t.Run("Unable to get IP", func(t *testing.T) {
		DefaultHTTPGetAddress = "http://127.0.0.1:12345"

		_, err := infoFunction.handler(context.TODO())
		if err == nil {
			t.Fatal("Error failed to trigger with an invalid request")
		}
	})

	t.Run("Non 200 Response", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
		}))
		defer ts.Close()

		DefaultHTTPGetAddress = ts.URL

		resp, err := infoFunction.handler(context.TODO())
		if err != nil && err.Error() != ErrNon200Response.Error() {
			t.Fatalf("Error failed to trigger with an invalid HTTP response: %v", err)
		}

		infoFunction.logger.Debug(resp)
	})

	t.Run("Unable decode IP", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
		}))
		defer ts.Close()

		DefaultHTTPGetAddress = ts.URL

		_, err := infoFunction.handler(context.TODO())
		if err == nil {
			t.Fatal("Error failed to trigger with an invalid HTTP response")
		}
	})

	t.Run("Successful Request", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			_, _ = fmt.Fprintf(w, "127.0.0.1")
		}))
		defer ts.Close()

		DefaultHTTPGetAddress = ts.URL

		response, err := infoFunction.handler(context.TODO())
		if err != nil {
			t.Fatal("Everything should be ok")
		}

		assert.Equal(t, response.StatusCode, 200)
		var infoE main.main.AppInfo
		_ = json.Unmarshal([]byte(response.Body), &infoE)
		assert.Equal(t, infoE.AppName, infoA.AppName)
		assert.Equal(t, infoE.Description, infoA.Description)
	})
}
