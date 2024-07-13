package blocks

import (
	"bytes"
	"encoding/gob"

	"github.com/fahyjo/blockchain/crypto"
	proto "github.com/fahyjo/blockchain/proto"
	"github.com/fahyjo/blockchain/transactions"
)

// EncodeBlock encodes a Block into a byte slice
func EncodeBlock(b *Block) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(b)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecodeBlock decodes a byte slice into a Block
func DecodeBlock(b []byte) (*Block, error) {
	var block Block
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err := dec.Decode(&block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

// EncodeHeader encodes a Header into a byte slice
func EncodeHeader(h *Header) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(h)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ConvertProtoBlock converts the given proto block into a domain Block
func ConvertProtoBlock(protoBlock *proto.Block) *Block {
	header := convertProtoHeader(protoBlock.Header)
	var txs []*transactions.Transaction
	for _, protoTx := range protoBlock.Transactions {
		txs = append(txs, transactions.ConvertProtoTransaction(protoTx))
	}
	sig := crypto.NewSignature(protoBlock.Sig)
	pubKey := crypto.NewPublicKey(protoBlock.PubKey)
	return NewBlock(header, txs, sig, pubKey)
}

// convertProtoHeader converts the given proto header into a domain Header
func convertProtoHeader(protoHeader *proto.Header) *Header {
	return NewHeader(protoHeader.Height, protoHeader.PrevHash, protoHeader.RootHash, protoHeader.TimeStamp)
}
