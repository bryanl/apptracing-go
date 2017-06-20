package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/bryanl/apptracing-go/internal/platform/logging"
	"github.com/bryanl/apptracing-go/internal/platform/tracing"
	"github.com/go-kit/kit/log"
	opentracing "github.com/opentracing/opentracing-go"
	olog "github.com/opentracing/opentracing-go/log"
)

func main() {
	rand.Seed(time.Now().Unix())
	logger := logging.Init("func")

	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	tracer, err := tracing.Init("goParallel", logger)
	if err != nil {
		logger.Log("msg", err)
		return
	}
	defer tracer.Close()

	span := opentracing.GlobalTracer().StartSpan("parallel")
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	delays := []int{45, 55, 65, 95, 105, 10, 15, 25, 30, 59, 93, 71, 9, 15, 35, 47}

	var wg sync.WaitGroup
	for i := range delays {
		delay := delays[i]
		go func(id, delay int) {
			wg.Add(1)
			defer wg.Done()
			task(ctx, fmt.Sprintf("task-%d", id), delay, logger)
		}(i, delay)
	}

	wg.Wait()

	time.Sleep(1 * time.Second)
}

func task(ctx context.Context, name string, delay int, logger log.Logger) {
	span, _ := opentracing.StartSpanFromContext(ctx, name)
	defer span.Finish()

	logger.Log("msg", name)

	span.LogFields(
		olog.Int("delay", delay))

	time.Sleep(time.Duration(delay) * time.Millisecond)
}
