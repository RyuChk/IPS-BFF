package converter

import (
	"bytes"
	"encoding/gob"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func ToIntArray[T Number](input []T) []int {
	buffer := make([]int, len(input))
	for i, v := range input {
		buffer[i] = int(v)
	}

	return buffer
}

func ToByte[T any](input T) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(input)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func FromByte[T any](input []byte, output *T) error {
	enc := gob.NewDecoder(bytes.NewReader(input))
	return enc.Decode(output)
}
