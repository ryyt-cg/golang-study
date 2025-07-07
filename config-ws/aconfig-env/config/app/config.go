package app

import (
	"fmt"
	"github.com/cristalhq/aconfig"
	"github.com/go-playground/validator/v10"
	"os"
	"strings"
)

// Config stores the application-wide configurations
var (
	Config   appConfig
	validate *validator.Validate
)

type appConfig struct {
	AppInfo  AppInfoConfig
	Database DatabaseConfig
	Server   ServerConfig
}

// Validate all config required values are populated.
func (config appConfig) Validate() error {
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
// Environment variables with the prefix "NGEN_" in their names are also read automatically.
func LoadConfig(configPath string) error {
	env := strings.ToUpper(os.Getenv("ENV"))
	configFile := configPath + "/" + getConfigFile(env) + ".json"

	var cfg appConfig
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		SkipFlags:    true,
		SkipDefaults: true,
		Files:        []string{configFile},
		//FileDecoders: map[string]aconfig.FileDecoder{
		//	".yaml": aconfigyaml.New(), // Register the YAML decoder
		//},
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
