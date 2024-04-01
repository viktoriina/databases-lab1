package utils

import (
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
)

func EncodeBinary(v any) ([]byte, error) {
	const invalidSize = -1
	size := binary.Size(v)
	if size == invalidSize {
		return nil, errors.New("invalid value size")
	}
	buf := bytes.NewBuffer(make([]byte, 0, size))
	if err := binary.Write(buf, binary.BigEndian, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
