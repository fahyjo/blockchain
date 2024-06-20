package blocks

import proto "github.com/fahyjo/blockchain/proto"

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

func convertProtoHeader(protoHeader *proto.Header) *Header {
	return NewHeader(protoHeader.Height, protoHeader.PrevHash, protoHeader.RootHash, protoHeader.TimeStamp)
}
