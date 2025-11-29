package app

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/spf13/viper"
)

// Config stores the application-wide configurations
var Config AppConfig

type AppConfig struct {
	AppInfo  AppInfoConfig
	Database DatabaseConfig
	//Okta     OktaConfig
}

// Validate
// Make all config required values are populated.
func (config AppConfig) Validate() error {

	if err := config.AppInfo.Validate(); err != nil {
		panic(err.Error())
	}
	if err := config.Database.Validate(); err != nil {
		panic(err.Error())
	}
	//if err := config.Okta.OAuth2.Validate(); err != nil {
	//	panic(err.Error())
	//}
	return validation.ValidateStruct(&config) //validation.Field(&config.DSN, validation.Required),
}

// LoadConfig
// loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be named as app.yaml.
// Environment variables with the prefix "NGEN_" in their names are also read automatically.
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	env := strings.ToUpper(os.Getenv("ENV"))
	v.SetConfigName(getConfigFile(env))

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	configFile, err := Asset("config/" + getConfigFile(env) + ".yaml")
	if err != nil {
		return fmt.Errorf("failed to read the configuration file in bindata: %s", err)
	}

	if err := v.ReadConfig(bytes.NewReader(configFile)); err != nil {
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
	case "DEV":
		return "app-dev"
	default:
		return "app"
	}
}
