package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/bryanl/apptracing-go/internal/platform/logging"
	atrand "github.com/bryanl/apptracing-go/internal/platform/rand"
	"github.com/go-kit/kit/log"
)

func main() {
	rand.Seed(time.Now().Unix())
	logger := logging.Init("func")

	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	ctx := context.Background()

	// initialize opentracing

	function1(ctx, logger)

	time.Sleep(1 * time.Second)
}

func function1(ctx context.Context, logger log.Logger) {
	logger.Log("msg", "function1")

	// extract span from context (if it exists)

	delay := atrand.Between(25, 50)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	function2(ctx, logger)
}

func function2(ctx context.Context, logger log.Logger) {
	logger.Log("msg", "function2")

	// extract span from context (if it exists)

	delay := atrand.Between(25, 50)
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
