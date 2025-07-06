package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"

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
	Info     InfoConfig
	Database DatabaseConfig
	Server   ServerConfig
}

// Validate all config required values are populated.
func (config appConfig) validate() error {
	validate = validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(config.Info); err != nil {
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

	// Load environment variables from .env file
	if err := loadEnv(); err != nil {
		return fmt.Errorf("fail to load environment variables: %s", err)
	}

	// TODO
	// will use reflect package to scan the Config struct
	// and replace any attributes's value that has prefix ${
	// and suffix } with the environment variable value
	//replaceWithEnv(&Config.Database.Postgres.Dsn)
	//replaceWithEnv(&Config.Database.Sqlite.Dsn)

	fields, _ := listStructKeys(Config)
	for _, field := range fields {
		if v.IsSet(field) {
			value := v.GetString(field)
			if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
				envVar := value[2 : len(value)-1]
				v.Set(field, v.GetString(envVar))
			}
		}
	}

	return Config.validate()
}

func listStructKeys(s interface{}) ([]string, error) {
	// Recursively get the config struct tag mapstructure
	keys := []string{}
	ct := reflect.TypeOf(s)

	if ct.Kind() != reflect.Struct {
		return nil, fmt.Errorf("listStructKeys: %v is not a struct", ct)
	}

	for i := range ct.NumField() {
		field := ct.Field(i)
		tag := field.Tag.Get("mapstructure")

		if field.Type.Kind() == reflect.Struct {
			res, err := listStructKeys(reflect.New(field.Type).Elem().Interface())
			if err != nil {
				return nil, err
			}
			for _, k := range res {
				keys = append(keys, fmt.Sprintf("%s.%s", tag, k))
			}
		} else {
			keys = append(keys, tag)
		}
	}

	return keys, nil
}

// loadEnv loads environment variables from the .env file.
func loadEnv() error {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("fail to read the .env file: %s", err)
	}

	return nil
}

// replaceConfigFile replaces the configuration file name based on the environment variable.
func replaceWithEnv(attr *string) {
	if strings.HasPrefix(*attr, "${") && strings.HasSuffix(*attr, "}") {
		envVar := (*attr)[2 : len(*attr)-1]
		*attr = viper.GetString(envVar)
	}

	//if strings.HasPrefix(Config.Database.Postgres.Dsn, "${") && strings.HasSuffix(Config.Database.Postgres.Dsn, "}") {
	//	envVar := Config.Database.Postgres.Dsn[2 : len(Config.Database.Postgres.Dsn)-1]
	//	Config.Database.Postgres.Dsn = viper.GetString(envVar)
	//}
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
