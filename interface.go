package typewriter

// Logger represents an interface that can be adapted by all loggers.
type Logger interface {
	// With appends the specified name to the name identifying the logger.
	With(name string) Logger

	// WithV specifies additional values to be logged in order to provide more context.
	WithV(values ...Value) Logger

	// Info logs the specified message and optionally a set of values.
	Info(message string, values ...Value)

	// Error logs the specified error together with a custom message and the logger's name and
	// context values.
	Error(err error, message string)
}

// Value represents an interface that converts an element to a string.
type Value interface {

	// String returns a string representation of the value to be printed.
	String() string
}
