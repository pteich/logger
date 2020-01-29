# logger
--
    import "github.com/pteich/logger"

## Usage

#### type Logger

```go
type Logger struct {
	zerolog.Logger
}
```

Logger represents a logger that embeds zerolog

#### func  New

```go
func New(opts ...Option) *Logger
```
New returns a new logger with specific options (zero to x)

#### func  NewFromZerolog

```go
func NewFromZerolog(logger zerolog.Logger) *Logger
```
NewFromZerolog returns a new logger from an existing zerolog instance

#### func (*Logger) Log

```go
func (log *Logger) Log(v ...interface{}) error
```
Log implements a common log interface

#### type Option

```go
type Option func(c *config)
```

Options is an config setting function

#### func  WithLogConsole

```go
func WithLogConsole() Option
```
WithLogConsole enabled human readable logging to console

#### func  WithLogLevel

```go
func WithLogLevel(logLevel string) Option
```
WithLogLevel sets a specific loglevel given as string (e.g. debug)

#### func  WithNanoSeconds

```go
func WithNanoSeconds() Option
```
WithNanoSeconds logs timestamp with nanos if millis are not enough

#### func  WithServiceName

```go
func WithServiceName(serviceName string) Option
```
WithServiceName sets a service name as log field
