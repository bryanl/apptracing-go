package tracing

import (
	"fmt"
	"io"

	"github.com/go-kit/kit/log"
	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics"
)

type tracingLogger struct {
	logger log.Logger
}

func (l *tracingLogger) Error(msg string) {
	_ = l.logger.Log("err", msg)
}

func (l *tracingLogger) Infof(msg string, args ...interface{}) {
	_ = l.logger.Log("msg", fmt.Sprintf(msg, args...))
}

// Init initializes opentracing.
func Init(serviceName string, logger log.Logger) (io.Closer, error) {
	al := &tracingLogger{logger: logger}

	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "localhost:5775",
		},
	}

	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	closer, err := cfg.InitGlobalTracer(
		serviceName,
		jaegercfg.Logger(al),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		logger.Log("msg", "could not initialize jaeger tracer", "err", err.Error())
		return nil, err
	}

	return closer, nil
}
