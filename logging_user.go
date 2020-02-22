package typewriter

import "fmt"

// UserLogger logs with a human-readable format for applications where logs are only informational
// and do not need to be imported into other programs.
type userLogger struct {
	name   string
	values []Value
}

// NewUserLogger returns a newly configured user logger with the given name.
func NewUserLogger(name string) Logger {
	return userLogger{name, make([]Value, 0)}
}

func (log userLogger) With(name string) Logger {
	return userLogger{
		name:   fmt.Sprintf("%s.%s", log.name, name),
		values: log.values,
	}
}

func (log userLogger) WithV(values ...Value) Logger {
	v := copyValueSlice(log.values)
	for _, value := range values {
		v = append(v, value)
	}
	return userLogger{
		name:   log.name,
		values: v,
	}
}

func (log userLogger) Info(message string, values ...Value) {
	str := fmt.Sprintf("[INFO] %s (%s)", log.name, timeNow())
	if kv := log.valueString(log.values); kv != "" {
		str += " " + kv
	}

	if len(values) == 0 {
		fmt.Printf("%s => %s\n", str, message)
		return
	}

	valStr := log.valueString(values)
	if valStr != "" {
		valStr = " [" + valStr + "]"
	}

	fmt.Printf("%s => %s%s\n", str, message, valStr)
}

func (log userLogger) Error(err error, message string) {
	str := fmt.Sprintf("[ERROR] %s (%s)", log.name, timeNow())
	if kv := log.valueString(log.values); kv != "" {
		str += " " + kv
	}
	fmt.Printf("%s => %s [%s]\n", str, message, err)
}

func (log userLogger) valueString(values []Value) string {
	if len(values) == 0 {
		return ""
	}

	res := ""
	for i, v := range values {
		if i == 0 {
			res += v.String()
		} else {
			res += " | " + v.String()
		}
	}
	return res
}
