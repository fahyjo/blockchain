package blocks

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
