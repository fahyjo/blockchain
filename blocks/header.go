package blocks

// Header contains metadata about a Block
type Header struct {
	Height    int64  // Height is the block index in the chain
	PrevHash  []byte // PrevHash is the hash of the block immediately preceding this Block
	RootHash  []byte // RootHash is the root hash of the merkle tree made up of this Block's transactions
	TimeStamp int64  // TimeStamp is the time this Block was created
}

// NewHeader creates a new Header struct
func NewHeader(height int64, prevHash []byte, rootHash []byte, timeStamp int64) *Header {
	return &Header{
		Height:    height,
		PrevHash:  prevHash,
		RootHash:  rootHash,
		TimeStamp: timeStamp,
	}
}
