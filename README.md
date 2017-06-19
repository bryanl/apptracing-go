# Velocity 2017: Go AppTracing Exercise

To fetch the source:

`go get -u github.com/bryanl/apptracing-go`

## Getting started

1. Create docker environment: `docker-compose up -d`
1. Create database: `make create-db`
1. Import test data: `make import-people`

Exercises:

* [Tracing Functions](functions)
* [Tracing HTTP client/server](clientserver)
* [Exploring a bigger app](app)