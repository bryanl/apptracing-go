package main

import (
	"context"
	"fmt"
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

	tracer, err := tracing.Init("goSerial", logger)
	if err != nil {
		logger.Log("msg", err)
		return
	}
	defer tracer.Close()

	span := opentracing.GlobalTracer().StartSpan("serial")
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	delays := []int{45, 55, 65, 95, 105, 10, 15, 25, 30, 59, 93, 71, 9, 15, 35, 47}
	for i, delay := range delays {
		task(ctx, fmt.Sprintf("task-%d", i), delay, logger)
	}

	time.Sleep(1 * time.Second)
}

func task(ctx context.Context, name string, delay int, logger log.Logger) {
	span, ctx := opentracing.StartSpanFromContext(ctx, name)
	defer span.Finish()

	logger.Log("msg", name)

	span.LogFields(
		olog.Int("delay", delay))

	time.Sleep(time.Duration(delay) * time.Millisecond)
}
