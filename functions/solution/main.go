package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/bryanl/apptracing/internal/platform/logging"
	"github.com/bryanl/apptracing/internal/platform/tracing"
	"github.com/go-kit/kit/log"
	opentracing "github.com/opentracing/opentracing-go"
	olog "github.com/opentracing/opentracing-go/log"
)

func main() {
	rand.Seed(time.Now().Unix())
	logger := logging.Init("func")

	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	tracer, err := tracing.Init("goFunc", logger)
	if err != nil {
		logger.Log("msg", err)
		return
	}
	defer tracer.Close()

	ctx := context.Background()
	function1(ctx, logger)

	time.Sleep(1 * time.Second)
}

func function1(ctx context.Context, logger log.Logger) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "function1")
	defer span.Finish()

	span.SetBaggageItem("token", "token")

	logger.Log("msg", "function1")

	delay := random(25, 50)
	span.LogFields(
		olog.Int("delay", delay))

	time.Sleep(time.Duration(delay) * time.Millisecond)

	function2(ctx, logger)
}

func function2(ctx context.Context, logger log.Logger) {
	span, _ := opentracing.StartSpanFromContext(ctx, "function2")
	defer span.Finish()

	token := span.BaggageItem("token")

	logger.Log("msg", "function2", "token", token)

	delay := random(25, 50)
	span.LogFields(
		olog.Int("delay", delay))
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
