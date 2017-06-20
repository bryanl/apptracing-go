package main

import (
	"io/ioutil"
	"net/http"

	"github.com/bryanl/apptracing-go/internal/platform/logging"
	"github.com/bryanl/apptracing-go/internal/platform/tracing"
	opentracing "github.com/opentracing/opentracing-go"
)

func main() {
	logger := logging.Init("client")

	tracer, err := tracing.Init("goClient", logger)
	if err != nil {
		logger.Log("msg", err)
		return
	}
	defer tracer.Close()

	span := opentracing.StartSpan("request")
	defer span.Finish()

	req, _ := http.NewRequest("GET", "http://localhost:8081", nil)

	opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Log("msg", "http client error", "err", err.Error())
	}

	if sc := resp.StatusCode; sc != http.StatusOK {
		logger.Log("msg", "unexpected http status code",
			"status-code", sc)
		return
	}

	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	logger.Log("msg", string(b))
}
