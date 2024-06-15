package blocks

import (
	"bytes"
	"encoding/gob"
)

type Header struct {
	Height    int64
	PrevHash  []byte
	RootHash  []byte
	TimeStamp int64
}

func NewHeader(height int64, prevHash []byte, rootHash []byte, timeStamp int64) *Header {
	return &Header{
		Height:    height,
		PrevHash:  prevHash,
		RootHash:  rootHash,
		TimeStamp: timeStamp,
	}
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
