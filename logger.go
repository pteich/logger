package logger

import (
	"io"
	"os"
	"runtime"
	"time"

	"github.com/rs/zerolog"
)

type Config struct {
	logLevel    string
	logConsole  bool
	serviceName string
}

// Logger represents a logger that embeds zerolog
type Logger struct {
	zerolog.Logger
}

// New returns a new logger with specific options
func New(opts ...Option) *Logger {

	_, filename, _, _ := runtime.Caller(0)
	config := Config{
		logLevel:    "debug",
		logConsole:  true,
		serviceName: filename,
	}

	for _, opt := range opts {
		opt(&config)
	}

	zerologLevel, err := zerolog.ParseLevel(config.logLevel)
	if err == nil {
		zerolog.SetGlobalLevel(zerologLevel)
	}

	var logDest io.Writer = os.Stdout

	if config.logConsole {
		logDest = zerolog.ConsoleWriter{Out: logDest, TimeFormat: time.RFC3339}
	}

	return &Logger{
		Logger: zerolog.New(logDest).With().Timestamp().Str("service", config.serviceName).Logger(),
	}
}

// Log implements a common log interface
func (log *Logger) Log(v ...interface{}) error {
	log.Print(v...)
	return nil
}

// NewFromZerolog returns a new logger from an existing zerolog instance
func NewFromZerolog(logger zerolog.Logger) *Logger {
	return &Logger{logger}
}
