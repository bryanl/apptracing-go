package main

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/bryanl/apptracing-go/internal/platform/logging"
	"github.com/bryanl/apptracing-go/internal/platform/tracing"
	"github.com/go-kit/kit/log"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	olog "github.com/opentracing/opentracing-go/log"
)

func main() {
	rand.Seed(time.Now().Unix())
	logger := logging.Init("func")

	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	tracer, err := tracing.Init("goTagging", logger)
	if err != nil {
		logger.Log("msg", err)
		return
	}
	defer tracer.Close()

	ctx := context.Background()

	delay := 95
	task(ctx, "tagger", delay, logger)

	time.Sleep(1 * time.Second)
}

func task(ctx context.Context, name string, delay int, logger log.Logger) {
	span, _ := opentracing.StartSpanFromContext(ctx, name)
	defer span.Finish()

	logger.Log("msg", name)

	span.SetTag("action", "task")
	err := errors.New("boom")
	ext.Error.Set(span, true)
	span.SetTag("error.object", err)

	span.LogFields(
		olog.Int("delay", delay))

	time.Sleep(time.Duration(delay) * time.Millisecond)
}
