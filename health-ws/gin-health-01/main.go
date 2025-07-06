package main

import (
	"github.com/gin-gonic/gin"
	healthcheck "github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	"github.com/tavsec/gin-healthcheck/config"
)

func main() {
	r := gin.Default()

	/*
		Example of a ping check that pings https://www.google.com with a GET request
			[
			  {
				"name": "ping-https://www.google.com",
				"pass": true
			  }
			]
	*/
	healthConfig := config.DefaultConfig()
	healthConfig.HealthPath = "/health"
	googleCheck := checks.NewPingCheck("https://www.google.com", "GET", 1000, nil, nil)
	msCheck := checks.NewPingCheck("https://www.microsoft.com", "GET", 1000, nil, nil)
	checks := []checks.Check{googleCheck, msCheck}

	healthcheck.New(r, healthConfig, checks)
	r.Run()
}
