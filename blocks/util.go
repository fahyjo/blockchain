package blocks

import (
	"bytes"
	"encoding/gob"
)

func EncodeBlock(b *Block) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(b)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecodeBlock(b []byte) (*Block, error) {
	var block Block
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err := dec.Decode(&block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func EncodeHeader(h *Header) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(h)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
