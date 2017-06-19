package main

import (
	"fmt"
	stdlog "log"
	"math/rand"
	"net/http"
	"time"

	"github.com/bryanl/apptracing/internal/platform/logging"
	"github.com/bryanl/apptracing/internal/platform/tracing"
	"github.com/go-kit/kit/log"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func main() {
	rand.Seed(time.Now().Unix())
	logger := logging.Init("server")

	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	tracer, err := tracing.Init("goServer", logger)
	if err != nil {
		logger.Log("msg", err)
		return
	}
	defer tracer.Close()

	http.Handle("/", handler(logger))
	stdlog.Fatal(http.ListenAndServe("localhost:8081", nil))
}

func handler(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trace := false
		wireContext, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))

		if err == nil {
			trace = true
		}

		var span opentracing.Span

		if trace {
			span = opentracing.StartSpan(
				"handler",
				ext.RPCServerOption(wireContext))
			defer span.Finish()
		}

		statusCode := 200

		if trace {
			ext.HTTPStatusCode.Set(span, uint16(statusCode))
			ext.HTTPMethod.Set(span, r.Method)
		}

		w.WriteHeader(statusCode)
		fmt.Fprintln(w, "hello world")
	}
}
