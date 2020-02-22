package typewriter

import (
	"time"
)

// V is a short typealias for V
type V = map[string]string

func copyValueSlice(old []Value) []Value {
	new := make([]Value, len(old))
	copy(new, old)
	return new
}

func timeNow() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05.000")
}
