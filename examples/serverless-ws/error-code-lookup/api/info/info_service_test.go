package info

import (
	"error-code-lookup/config/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getAppInfo(t *testing.T) {
	// instantiate application config object
	app.Config = app.AppConfig{
		AppInfo: app.AppInfoConfig{
			Name:        "Error Code Lookup",
			Description: "Service to lookup error code definitions",
			Version:     "1.0.0",
		},
	}

	appInfo := NewService()
	result, _ := appInfo.getAppInfo()

	assert.Equal(t, "Error Code Lookup", result.AppName)
	assert.Equal(t, "1.0.0", result.Version)
	assert.Equal(t, "Service to lookup error code definitions", result.Description)
}
