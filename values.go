package typewriter

// KV represents a value that consists of a key and a value.
type KV struct {
	Key   string
	Value string
}

func (kv KV) String() string {
	return kv.Key + "=" + kv.Value
}

// K represents a value that merely consists of a key.
type K struct {
	Key string
}

func (k K) String() string {
	return k.Key
}
