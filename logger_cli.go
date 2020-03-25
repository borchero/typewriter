package typewriter

import (
	"github.com/fatih/color"
)

type cliLogger struct {
}

// NewCLILogger returns a newly configured logger that prints colored logs for CLI tools to enable
// capturing relevant logs quickly. The CLI logger does not provide any context.
func NewCLILogger() Logger {
	return cliLogger{}
}

func (log cliLogger) With(name string) Logger {
	return log
}

func (log cliLogger) WithV(values ...Value) Logger {
	return log
}

func (log cliLogger) Info(message string, values ...Value) {
	col := color.New(color.FgBlue).Add(color.Bold)

	valStr := log.valueString(values)
	if valStr != "" {
		valStr = " [" + valStr + "]"
	}

	col.Printf("%s%s\n", message, valStr)
}

func (log cliLogger) Error(err error, message string) {
	col := color.New(color.FgRed).Add(color.Bold)
	col.Printf("%s: %s\n", message, err)
}

func (cliLogger) valueString(values []Value) string {
	if len(values) == 0 {
		return ""
	}

	res := ""
	for i, v := range values {
		if i == 0 {
			res += v.String()
		} else {
			res += ", " + v.String()
		}
	}

	return res
}
