# Fiber 01 Example
Build a RESTful API using Fiber, a fast web framework for Go. This example demonstrates how to set up a basic application with routing, middleware, and error handling.

## Features
- Middleware
  - Accepts JSON 
  - Health Check
  - Monitor
  - Prometheus
  - Recovery
  - RequestID
  
## Endpoints
- `GET /api/gof/health`: Health check endpoint
- `GET /api/gof/metrics`: Prometheus metrics endpoint
- `GET /api/gof/monitor`: Monitor Page
- `GET /api/gof/info`: Application information endpoint
- `GET /api/gof/v1/authors/:id`: Get author by ID
- `GET /api/gof/v1/authors/names/:name`: Get author by name
- `GET /api/gof/v1/authors/?name=:name`: Get author by name with query parameter

## Load Testing
- Using Grafana k6 for load testing.

