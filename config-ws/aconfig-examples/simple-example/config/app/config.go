package app

import (
	"fmt"
	"github.com/cristalhq/aconfig"
	"github.com/go-playground/validator/v10"
	"os"
	"strings"
)

var (
	// Config stores the application-wide configurations
	Config AppConfig
	// validator is used to validate the configuration values
	// only visible within this package.
	validate *validator.Validate
)

type AppConfig struct {
	AppInfo  AppInfoConfig  `json:"appInfo" validate:"required"`
	Database DatabaseConfig `json:"database" validate:"required"`
	Server   ServerConfig   `json:"server" validate:"required"`
}

// Validate all config required values are populated.
func (config AppConfig) Validate() error {
	validate = validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(config.AppInfo); err != nil {
		panic(err.Error())
	}
	if err := validate.Struct(config.Database); err != nil {
		panic(err.Error())
	}

	return nil
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be named as app.yaml.
func LoadConfig(configPath string) error {
	env := strings.ToUpper(os.Getenv("ENV"))
	configFile := configPath + "/" + getConfigFile(env) + ".json"

	loader := aconfig.LoaderFor(&Config, aconfig.Config{
		SkipFlags: true,
		Files:     []string{configFile},
	})
	if err := loader.Load(); err != nil {
		return fmt.Errorf("failed to load configuration file %s: %w", configFile, err)
	}

	return Config.Validate()
}

func getConfigFile(env string) string {
	switch env {
	case "PRD":
		return "app-prd"
	case "DEV":
		return "app-dev"
	default:
		return "app"
	}
}
