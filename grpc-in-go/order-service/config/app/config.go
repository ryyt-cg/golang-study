package app

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// Config stores the application-wide configurations
var Config appConfig

type appConfig struct {
	Database DatabaseConfig
}

// Validate make sure all config required values are populated.
func (config appConfig) Validate() error {

	if err := config.Database.Validate(); err != nil {
		panic(err.Error())
	}

	return validation.ValidateStruct(&config) //validation.Field(&config.DSN, validation.Required),
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be named as app.yaml.
// Environment variables with the prefix "NGEN_" in their names are also read automatically.
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	env := strings.ToUpper(os.Getenv("ENV"))
	v.SetConfigName(getConfigFile(env))
	v.SetDefault("error_file", "config/errors.yaml")

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}

	return Config.Validate()
}

func getConfigFile(env string) string {
	switch env {
	case "PROD":
		return "app-prod"
	case "STG":
		return "app-stg"
	case "DEV":
		return "app-dev"
	default:
		return "app"
	}
}
