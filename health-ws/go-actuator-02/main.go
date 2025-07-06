package main

import (
	"github.com/gin-gonic/gin"
	actuator "github.com/sinhashubham95/go-actuator"
)

func main() {
	// create the gin engine
	engine := gin.Default()

	config := &actuator.Config{
		Endpoints: []int{
			actuator.Env,
			actuator.Info,
			actuator.Metrics,
		},
		Env:     "dev",
		Name:    "Naruto Rocks",
		Port:    8080,
		Version: "0.1.0",
	}

	// get the handler for actuator
	actuatorHandler := actuator.GetActuatorHandler(config)
	ginActuatorHandler := func(ctx *gin.Context) {
		actuatorHandler(ctx.Writer, ctx.Request)
	}

	engine.GET("/actuator/*endpoints", ginActuatorHandler)
	engine.Run()
}
