package logger

// Options is an config setting function
type Option func(c *Config)

// WithLogLevel sets a specific loglevel given as string (e.g. debug)
func WithLogLevel(logLevel string) Option {
	return func(c *Config) {
		c.logLevel = logLevel
	}
}

// WithLogConsole enabled human readable logging to console
func WithLogConsole(logConsole bool) Option {
	return func(c *Config) {
		c.logConsole = logConsole
	}
}

// WithServiceName sets a service name as log field
func WithServiceName(serviceName string) Option {
	return func(c *Config) {
		c.serviceName = serviceName
	}
}
