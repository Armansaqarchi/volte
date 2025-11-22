package utils

import (
	"bytes"
	"encoding/gob"
)

func GOBDecode(data []byte) (interface{}, error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	var out interface{}
	if err := dec.Decode(&out); err != nil {
		return nil, err
	}

	return out, nil
}

func GOBEncode(m map[string]int) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(m)
	return buf.Bytes(), err
}
