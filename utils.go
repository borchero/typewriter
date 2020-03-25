package typewriter

import (
	"fmt"
	"time"
)

// V returns a log-compatible instance of a string.
func V(value string) fmt.Stringer {
	return k{value}
}

// I returns a log-compatible instance of any interface. This method might be slow.
func I(value interface{}) fmt.Stringer {
	return k{fmt.Sprint(value)}
}

// KV returns a log-compatible mapping from a string key to a string value.
func KV(key, value string) fmt.Stringer {
	return kv{key, value}
}

// KI returns a log-compatible mapping from a string key to any value. This method might be slow.
func KI(key string, value interface{}) fmt.Stringer {
	return kv{key, fmt.Sprint(value)}
}

type k struct {
	key string
}

type kv struct {
	key   string
	value string
}

func (v k) String() string {
	return v.key
}

func (v kv) String() string {
	return v.key + "=" + v.value
}

func copySlice(old []fmt.Stringer) []fmt.Stringer {
	new := make([]fmt.Stringer, len(old))
	copy(new, old)
	return new
}

func concatenate(values []fmt.Stringer, delimiter, start, end string) string {
	if len(values) == 0 {
		return ""
	}

	res := ""
	for i, v := range values {
		if i == 0 {
			res += v.String()
		} else {
			res += delimiter + v.String()
		}
	}

	return start + res + end
}

func timeNow() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05.000")
}
