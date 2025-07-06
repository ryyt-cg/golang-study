# Health Check
Health Check is a simple endpoint that returns the health status of the service.  The health status allows Orchestrator to either restart the service or take it out of the load balancer pool if it is unhealthy.

There are many health check libraries available for Go.  In this module, I will use  [Gin Health Check](https://github.com/tavsec/gin-healthcheck).

