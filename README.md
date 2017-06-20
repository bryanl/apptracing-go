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



```
docker run -it --network apptracinggo_velocity --rm -v $(pwd)/assets:/assets postgres:9.6 bash
```

```
createdb -h apptracinggo_db_1 -U postgres velocity2017
psql -h apptracinggo_db_1  -U postgres -f /assets/import_people.sql
```
