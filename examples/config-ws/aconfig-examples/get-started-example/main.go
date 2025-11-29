package main

import (
	"fmt"
	"log"

	"github.com/cristalhq/aconfig"
)

type AppConfig struct {
	HTTPPort int `default:"1111" usage:"just give a number"`
	Auth     struct {
		User string `default:"def-user" usage:"your user"`
		Pass string `default:"def-pass" usage:"make it strong"`
	}
	Pass string `default:"" env:"SECRET" flag:"sec_ret"`
}

func main() {
	var cfg AppConfig
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		EnvPrefix:  "APP",
		FlagPrefix: "app",
		Files:      []string{"config/app.json"},
	})

	if err := loader.Load(); err != nil {
		log.Panic(err)
	}

	fmt.Printf("HTTPPort:  %v\n", cfg.HTTPPort)
	fmt.Printf("Auth.User: %q\n", cfg.Auth.User)
	fmt.Printf("Auth.Pass: %q\n", cfg.Auth.Pass)

}
