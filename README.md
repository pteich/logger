# logger
--
    import "github.com/pteich/logger"


## Usage

#### type Config

```go
type Config struct {
}
```


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
New returns a new logger with specific options

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
type Option func(c *Config)
```

Options is an config setting function

#### func  WithLogConsole

```go
func WithLogConsole(logConsole bool) Option
```
WithLogConsole enabled human readable logging to console

#### func  WithLogLevel

```go
func WithLogLevel(logLevel string) Option
```
WithLogLevel sets a specific loglevel given as string (e.g. debug)

#### func  WithServiceName

```go
func WithServiceName(serviceName string) Option
```
WithServiceName sets a service name as log field
