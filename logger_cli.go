package typewriter

import (
	"fmt"

	"github.com/fatih/color"
)

// CLILogger provides a logger that prints colored logs for CLI tools to enable users to capture
// relevant logs quickly. As opposed to other loggers, this logger does neither carry a name, nor
// provides any context. Note that this logger is generally slower than other loggers.
type CLILogger struct {
}

// NewCLILogger returns a newly configured CLILogger.
func NewCLILogger() CLILogger {
	return CLILogger{}
}

// With is a no-op which returns the CLI logger that the method is called on.
func (log CLILogger) With(name string) Logger {
	return log
}

// WithV is a no-op which returns the CLI logger that the method is called on.
func (log CLILogger) WithV(value fmt.Stringer, others ...fmt.Stringer) Logger {
	return log
}

// Info logs the specified message along with a set of values. Logging color will be blue. A newline
// will be added automatically.
func (log CLILogger) Info(message string, values ...fmt.Stringer) {
	col := color.New(color.FgBlue).Add(color.Bold)
	col.Printf("%s%s\n", message, concatenate(values, ", ", " [", "]"))
}

// Infof logs a format string. Logging color will be blue. A newline will be added automatically.
func (log CLILogger) Infof(format string, values ...interface{}) {
	col := color.New(color.FgBlue).Add(color.Bold)
	col.Printf(format+"\n", values)
}

// Error logs the specified error and message along with the given values. Logging color will be
// red. A newline will be added automatically.
func (log CLILogger) Error(message string, err error, values ...fmt.Stringer) {
	col := color.New(color.FgRed).Add(color.Bold)
	col.Printf("%s: %s%s\n", message, err, concatenate(values, ", ", " [", "]"))
}

// Errorf logs a format string. Logging color will be red. A newline will be added automatically.
func (log CLILogger) Errorf(format string, values ...interface{}) {
	col := color.New(color.FgRed).Add(color.Bold)
	col.Printf(format+"\n", values)
}

// Success logs a success message. Logging color will be green, a newline is added automatically.
func (log CLILogger) Success(message string) {
	col := color.New(color.FgGreen).Add(color.Bold)
	col.Printf("%s\n", message)
}
