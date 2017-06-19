package logging

import (
	stdlog "log"
	"os"

	"github.com/go-kit/kit/log"
)

// Init initializes the logger.
func Init(name string) log.Logger {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	logger = log.With(logger, "app", name)

	// send stdlib logging to go kit
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	stdlog.SetFlags(0)

	return logger
}
