package logger

// Options is an config setting function
type Option func(c *config)

// WithLogLevel sets a specific loglevel given as string (e.g. debug)
func WithLogLevel(logLevel string) Option {
	return func(c *config) {
		c.logLevel = logLevel
	}
}

// WithLogConsole enabled human readable logging to console
func WithLogConsole() Option {
	return func(c *config) {
		c.logConsole = true
	}
}

// WithNanoSeconds logs timestamp with nanos if millis are not enough
func WithNanoSeconds() Option {
	return func(c *config) {
		c.nanoSeconds = true
	}
}

// WithServiceName sets a service name as log field
func WithServiceName(serviceName string) Option {
	return func(c *config) {
		c.serviceName = serviceName
	}
}
