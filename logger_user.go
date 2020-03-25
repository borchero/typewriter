package typewriter

import "fmt"

// UserLogger provides a logger that prints logs for long-running processes (e.g. servers) in a
// human-readable format. This format is not necessarily suitable for tools analyzing logs.
type UserLogger struct {
	name   string
	values []fmt.Stringer
}

// NewUserLogger returns a newly configured UserLogger with the specified name.
func NewUserLogger(name string) Logger {
	return UserLogger{name, make([]fmt.Stringer, 0)}
}

// With appends the specified name to the logger's name, seperated by a period character.
func (log UserLogger) With(name string) Logger {
	return UserLogger{
		name:   fmt.Sprintf("%s.%s", log.name, name),
		values: log.values,
	}
}

// WithV specifies additional values to be logged whenever the logger is logging.
func (log UserLogger) WithV(value fmt.Stringer, values ...fmt.Stringer) Logger {
	v := copySlice(log.values)
	for _, value := range values {
		v = append(v, value)
	}
	return UserLogger{
		name:   log.name,
		values: v,
	}
}

// Info logs the specified message along with the given values. It also adds the logging level and
// a timestamp.
func (log UserLogger) Info(message string, values ...fmt.Stringer) {
	str := fmt.Sprintf("[INFO] %s (%s)", log.name, timeNow())
	str += concatenate(log.values, " | ", " [", "]")

	if len(values) == 0 {
		fmt.Printf("%s => %s\n", str, message)
		return
	}

	message += concatenate(values, " | ", " [", "]")
	fmt.Printf("%s => %s\n", str, message)
}

func (log UserLogger) Error(message string, err error, values ...fmt.Stringer) {
	str := fmt.Sprintf("[ERROR] %s (%s)", log.name, timeNow())
	str += concatenate(log.values, " | ", " [", "]")

	suffix := concatenate(log.values, " | ", " [", "]")
	if err != nil {
		suffix = ": " + fmt.Sprint(err) + suffix
	}

	fmt.Printf("%s => %s%s\n", str, message, suffix)
}
