package main

import (
	"context"
	"flag"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// db driver

	"github.com/bryanl/apptracing-go/internal/platform/logging"
	_ "github.com/lib/pq"
)

var ()

func main() {
	var (
		httpAddr = flag.String("http.addr", ":9999", "HTTP address")
	)

	rand.Seed(time.Now().Unix())

	logger := logging.Init("people")
	logger.Log("msg", "initializing")
	defer logger.Log("msg", "goodbye")

	ctx, cancel := context.WithCancel(context.Background())

	mux := initHTTP()
	srv := http.Server{Addr: *httpAddr, Handler: mux}
	logger.Log("addr", *httpAddr, "msg", "starting http server")

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.Log("msg", "http server error",
				"err", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	logger.Log("msg", "shutting down")
	cancel()

	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log("msg", "unable to stop http server",
			"err", err)
	}
}
