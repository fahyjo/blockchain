package blocks

import (
	"bytes"
	"encoding/gob"

	"github.com/fahyjo/blockchain/crypto"
	proto "github.com/fahyjo/blockchain/proto"
	"github.com/fahyjo/blockchain/transactions"
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
