package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

const (
	LogFieldService  = "service"
	LogFieldHostname = "hostname"
)

var StdLogger = New()

type config struct {
	level         string
	consoleOutput bool
	color         bool
	serviceName   string
	nanoSeconds   bool
	hostName      string
	caller        bool
}

// Logger represents a logger that embeds zerolog
type Logger struct {
	zerolog.Logger
}

// New returns a new logger with specific options (zero to x)
func New(opts ...Option) *Logger {
	config := config{
		level:         "debug",
		consoleOutput: false,
		serviceName:   "",
		nanoSeconds:   false,
		hostName:      "",
		color:         false,
	}

	for _, opt := range opts {
		opt(&config)
	}

	zerologLevel, err := zerolog.ParseLevel(config.level)
	if err == nil {
		zerolog.SetGlobalLevel(zerologLevel)
	}

	var logDest io.Writer = os.Stdout

	if config.consoleOutput {
		logDest = zerolog.ConsoleWriter{Out: logDest, TimeFormat: time.RFC3339, NoColor: !config.color}
	}

	// this is actually bad b/c we set package global variables in zerolog
	// but this is the only way (that I found) to do this
	if config.nanoSeconds {
		zerolog.DurationFieldInteger = true
		zerolog.DurationFieldUnit = time.Millisecond
		zerolog.TimeFieldFormat = time.RFC3339Nano
	}

	loggerContext := zerolog.New(logDest).Level(zerologLevel).With().Timestamp()
	if config.hostName != "" {
		loggerContext = loggerContext.Str(LogFieldHostname, config.hostName)
	}
	if config.serviceName != "" {
		loggerContext = loggerContext.Str(LogFieldService, config.serviceName)
	}
	if config.caller {
		loggerContext = loggerContext.Caller()
	}

	return &Logger{
		Logger: loggerContext.Logger(),
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
