package typewriter

import (
	"fmt"
	"os"
)

// Logger represents an interface that can be adapted by all loggers. Specific loggers may expose
// more functionality.
type Logger interface {

	// With appends the specified name to the name identifying the logger.
	With(name string) Logger

	// WithV specifies additional values to be logged in order to provide more context.
	WithV(value fmt.Stringer, others ...fmt.Stringer) Logger

	// Info logs the specified message and optionally a set of values.
	Info(message string, values ...fmt.Stringer)

	// Error logs the specified error together with a custom message and the logger's name and
	// context values. Optionally, a set of values is logged as well. The error may be nil.
	Error(message string, err error, values ...fmt.Stringer)
}

// Fail calls the Error method of the given logger and subsequently exits the program with error
// code 1.
func Fail(logger Logger, message string, err error) {
	logger.Error(message, err)
	os.Exit(1)
}
