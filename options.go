package logger

// Options is an config setting function
type Option func(c *config)

// WithLogLevel sets a specific loglevel given as string (e.g. debug)
func WithLogLevel(logLevel string) Option {
	return func(c *config) {
		c.level = logLevel
	}
}

// WithLogConsole enables human readable logging to console
func WithLogConsole(logConsole bool) Option {
	return func(c *config) {
		c.consoleOutput = logConsole
	}
}

// WithColor enables colored output for console logs
func WithColor(color bool) Option {
	return func(c *config) {
		c.color = color
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

// WithHostname sets a host name as log field
func WithHostname(hostName string) Option {
	return func(c *config) {
		c.hostName = hostName
	}
}

// WithCaller adds a field with the caller (Go file name and line)
func WithCaller() Option {
	return func(c *config) {
		c.caller = true
	}
}
