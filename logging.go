package tw

import (
	"fmt"
	"time"
)

// V is a short typealias for V
type V = map[string]string

// Logger represents an interface that can be adapted by all loggers.
type Logger interface {
	// WithName appends the specified name to the name identifying the logger.
	WithName(name string) Logger

	// WithKeyValue specifies and additional key-value pair to be logged to provide more context.
	WithKeyValue(key, value string) Logger

	// Info logs the specified message together with the logger's name and context KV pairs.
	Info(message string)

	// InfoKeyValue logs the specified message and a set of values.
	InfoKeyValue(message string, values V)

	// Error logs the specified error together with a custom message and the logger's name and
	// context KV pairs.
	Error(err error, message string)
}

// UserLogger logs with a human-readable format for applications where logs are only informational
// and do not need to be imported into other programs.
type userLogger struct {
	name string
	keys []string
	kv   V
}

// NewUserLogger returns a newly configured user logger with the given name.
func NewUserLogger(name string) Logger {
	return userLogger{name, make([]string, 0), make(V)}
}

func (log userLogger) WithName(name string) Logger {
	return userLogger{
		name: fmt.Sprintf("%s.%s", log.name, name),
		keys: log.keys,
		kv:   log.kv,
	}
}

func (log userLogger) WithKeyValue(key, value string) Logger {
	m := copyMap(log.kv)
	m[key] = value
	return userLogger{
		name: log.name,
		keys: append(copySlice(log.keys), key),
		kv:   m,
	}
}

func (log userLogger) Info(message string) {
	log.InfoKeyValue(message, V{})
}

func (log userLogger) InfoKeyValue(message string, values V) {
	str := fmt.Sprintf("[INFO] %s (%s)", log.name, timeNow())
	if kv := log.kvString(); kv != "" {
		str += " " + kv
	}

	if len(values) == 0 {
		fmt.Printf("%s => %s\n", str, message)
		return
	}

	kvs := ""
	i := 0
	for k, v := range values {
		if i == 0 {
			kvs += fmt.Sprintf("%s=%s", k, v)
		} else {
			kvs += fmt.Sprintf(" | %s=%s", k, v)
		}
		i++
	}
	fmt.Printf("%s => [%s] %s\n", str, kvs, message)
}

func (log userLogger) Error(err error, message string) {
	str := fmt.Sprintf("[ERROR] %s (%s)", log.name, timeNow())
	if kv := log.kvString(); kv != "" {
		str += " " + kv
	}
	fmt.Printf("%s => [%s] %s\n", str, message, err)
}

func (log userLogger) kvString() string {
	if len(log.keys) == 0 {
		return ""
	}
	res := ""
	for i, k := range log.keys {
		if i == 0 {
			res += fmt.Sprintf("%s=%s", k, log.kv[k])
		} else {
			res += fmt.Sprintf(" | %s=%s", k, log.kv[k])
		}
	}
	return res
}

func copySlice(old []string) []string {
	new := make([]string, len(old))
	copy(new, old)
	return new
}

func copyMap(old V) V {
	new := make(V)
	for k, v := range old {
		new[k] = v
	}
	return new
}

func timeNow() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05.000")
}
