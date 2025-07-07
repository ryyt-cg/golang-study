package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cristalhq/aconfig"
)

type MyConfig struct {
	HTTPPort int `default:"1111" usage:"just a number" json:"httpPort"`
	Auth     struct {
		User string `default:"def-user" usage:"your user" json:"user"`
		Pass string `default:"def-pass" usage:"make it strong"`
	}
}

// Load defaults from struct definition and overwrite with a file. And then overwrite with environment variables.
func main() {
	//os.Setenv("EXAMPLE_HTTP_PORT", "3333")
	//os.Setenv("EXAMPLE_AUTH_USER", "env-user")
	//os.Setenv("EXAMPLE_AUTH_PASS", "env-pass")
	defer os.Clearenv()

	var cfg MyConfig
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		SkipFlags: true,
		EnvPrefix: "EXAMPLE",
		Files:     []string{"testdata/example_config.json"},
	})
	if err := loader.Load(); err != nil {
		log.Panic(err)
	}

	fmt.Printf("HTTPPort:  %v\n", cfg.HTTPPort)
	fmt.Printf("Auth.User: %v\n", cfg.Auth.User)
	fmt.Printf("Auth.Pass: %v\n", cfg.Auth.Pass)

}
