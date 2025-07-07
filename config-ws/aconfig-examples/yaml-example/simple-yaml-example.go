package main

import (
	"fmt"
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigyaml"
	"log"
	"os"
)

type MyConfig struct {
	HTTPPort int `yaml:"httpPort"`
	Auth     struct {
		User string `yaml:"user" validate:"required"` // Default value if not set
		Pass string `yaml:"pass" validate:"required"` // Default value if not set
	}
}

func main() {
	//os.Setenv("EXAMPLE_HTTP_PORT", "3333")
	os.Setenv("EXAMPLE_AUTH_USER", "env-user")
	//os.Setenv("EXAMPLE_AUTH_PASS", "env-pass")
	defer os.Clearenv()

	var cfg MyConfig
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		SkipFlags: true,
		EnvPrefix: "EXAMPLE",
		Files:     []string{"config/example_config.yaml"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": aconfigyaml.New(), // Register the YAML decoder
		},
	})
	if err := loader.Load(); err != nil {
		log.Panic(err)
	}

	fmt.Printf("HTTPPort:  %v\n", cfg.HTTPPort)
	fmt.Printf("Auth.User: %v\n", cfg.Auth.User)
	fmt.Printf("Auth.Pass: %v\n", cfg.Auth.Pass)

}
