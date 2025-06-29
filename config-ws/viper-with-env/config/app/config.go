package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	//_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
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
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	env := strings.ToUpper(os.Getenv("ENV"))
	v.SetConfigName(getConfigFile(env))

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("fail to read the configuration file: %s", err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}

	if err := loadEnv(); err != nil {
		return fmt.Errorf("fail to load environment variables: %s", err)
	}

	return Config.Validate()
}

// loadEnv loads environment variables from the .env file.
func loadEnv() error {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("fail to read the .env file: %s", err)
	}

	Config.Database.Sqlite.Dsn = viper.GetString("SQLITE_DSN")
	Config.Database.Postgres.Dsn = viper.GetString("POSTGRES_DSN")
	return nil
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
