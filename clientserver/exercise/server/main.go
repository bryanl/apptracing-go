package main

import (
	"fmt"
	stdlog "log"
	"net/http"

	"github.com/bryanl/apptracing-go/internal/platform/logging"
	"github.com/go-kit/kit/log"
)

func main() {
	logger := logging.Init("server")

	http.Handle("/", handler(logger))
	stdlog.Fatal(http.ListenAndServe("localhost:8081", nil))
}

func handler(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCode := 200

		w.WriteHeader(statusCode)
		fmt.Fprintln(w, "hello world")
	}
}
