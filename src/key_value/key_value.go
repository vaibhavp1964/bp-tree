package keyvalue

import "bytes"

// KeyValue ...
type KeyValue struct {
	Key   []byte
	Value []byte
}

// GetKeyValue ...
func GetKeyValue(key, value string) KeyValue {
	return KeyValue{
		Key:   []byte(key),
		Value: []byte(value),
	}
}

// HasSameKey ...
func (kv KeyValue) HasSameKey(key []byte) bool {
	return bytes.Compare(kv.Key, key) == 0
}
