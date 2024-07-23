// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: proto/blockchain.proto

package blockchain

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StringDump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int64    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Ids  []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *StringDump) Reset() {
	*x = StringDump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringDump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringDump) ProtoMessage() {}

func (x *StringDump) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringDump.ProtoReflect.Descriptor instead.
func (*StringDump) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{0}
}

func (x *StringDump) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *StringDump) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type BytesDump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int64    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Ids  [][]byte `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *BytesDump) Reset() {
	*x = BytesDump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesDump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesDump) ProtoMessage() {}

func (x *BytesDump) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesDump.ProtoReflect.Descriptor instead.
func (*BytesDump) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{1}
}

func (x *BytesDump) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *BytesDump) GetIds() [][]byte {
	if x != nil {
		return x.Ids
	}
	return nil
}

type Handshake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListenAddr      string   `protobuf:"bytes,1,opt,name=listenAddr,proto3" json:"listenAddr,omitempty"`
	PeerListenAddrs []string `protobuf:"bytes,2,rep,name=peerListenAddrs,proto3" json:"peerListenAddrs,omitempty"`
}

func (x *Handshake) Reset() {
	*x = Handshake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handshake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handshake) ProtoMessage() {}

func (x *Handshake) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handshake.ProtoReflect.Descriptor instead.
func (*Handshake) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{2}
}

func (x *Handshake) GetListenAddr() string {
	if x != nil {
		return x.ListenAddr
	}
	return ""
}

func (x *Handshake) GetPeerListenAddrs() []string {
	if x != nil {
		return x.PeerListenAddrs
	}
	return nil
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{3}
}

type PreVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockID []byte `protobuf:"bytes,1,opt,name=blockID,proto3" json:"blockID,omitempty"`
	Sig     []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	PubKey  []byte `protobuf:"bytes,3,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
}

func (x *PreVote) Reset() {
	*x = PreVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreVote) ProtoMessage() {}

func (x *PreVote) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreVote.ProtoReflect.Descriptor instead.
func (*PreVote) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{4}
}

func (x *PreVote) GetBlockID() []byte {
	if x != nil {
		return x.BlockID
	}
	return nil
}

func (x *PreVote) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *PreVote) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

type PreCommit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockID []byte `protobuf:"bytes,1,opt,name=blockID,proto3" json:"blockID,omitempty"`
	Sig     []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	PubKey  []byte `protobuf:"bytes,3,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
}

func (x *PreCommit) Reset() {
	*x = PreCommit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreCommit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreCommit) ProtoMessage() {}

func (x *PreCommit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreCommit.ProtoReflect.Descriptor instead.
func (*PreCommit) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{5}
}

func (x *PreCommit) GetBlockID() []byte {
	if x != nil {
		return x.BlockID
	}
	return nil
}

func (x *PreCommit) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *PreCommit) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *Header        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	Sig          []byte         `protobuf:"bytes,3,opt,name=sig,proto3" json:"sig,omitempty"`
	PubKey       []byte         `protobuf:"bytes,4,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{6}
}

func (x *Block) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *Block) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *Block) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height    int64  `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	PrevHash  []byte `protobuf:"bytes,2,opt,name=prevHash,proto3" json:"prevHash,omitempty"`
	RootHash  []byte `protobuf:"bytes,3,opt,name=rootHash,proto3" json:"rootHash,omitempty"`
	TimeStamp int64  `protobuf:"varint,4,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{7}
}

func (x *Header) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Header) GetPrevHash() []byte {
	if x != nil {
		return x.PrevHash
	}
	return nil
}

func (x *Header) GetRootHash() []byte {
	if x != nil {
		return x.RootHash
	}
	return nil
}

func (x *Header) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inputs  []*TxInput  `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs []*TxOutput `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{8}
}

func (x *Transaction) GetInputs() []*TxInput {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *Transaction) GetOutputs() []*TxOutput {
	if x != nil {
		return x.Outputs
	}
	return nil
}

type TxInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxID            []byte           `protobuf:"bytes,1,opt,name=txID,proto3" json:"txID,omitempty"`
	UtxoIndex       int64            `protobuf:"varint,2,opt,name=utxoIndex,proto3" json:"utxoIndex,omitempty"`
	UnlockingScript *UnlockingScript `protobuf:"bytes,3,opt,name=unlockingScript,proto3" json:"unlockingScript,omitempty"`
}

func (x *TxInput) Reset() {
	*x = TxInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxInput) ProtoMessage() {}

func (x *TxInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxInput.ProtoReflect.Descriptor instead.
func (*TxInput) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{9}
}

func (x *TxInput) GetTxID() []byte {
	if x != nil {
		return x.TxID
	}
	return nil
}

func (x *TxInput) GetUtxoIndex() int64 {
	if x != nil {
		return x.UtxoIndex
	}
	return 0
}

func (x *TxInput) GetUnlockingScript() *UnlockingScript {
	if x != nil {
		return x.UnlockingScript
	}
	return nil
}

type UnlockingScript struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKey []byte `protobuf:"bytes,1,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Sig    []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (x *UnlockingScript) Reset() {
	*x = UnlockingScript{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockingScript) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockingScript) ProtoMessage() {}

func (x *UnlockingScript) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockingScript.ProtoReflect.Descriptor instead.
func (*UnlockingScript) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{10}
}

func (x *UnlockingScript) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *UnlockingScript) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

type TxOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount        int64          `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	LockingScript *LockingScript `protobuf:"bytes,2,opt,name=lockingScript,proto3" json:"lockingScript,omitempty"`
}

func (x *TxOutput) Reset() {
	*x = TxOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxOutput) ProtoMessage() {}

func (x *TxOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxOutput.ProtoReflect.Descriptor instead.
func (*TxOutput) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{11}
}

func (x *TxOutput) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TxOutput) GetLockingScript() *LockingScript {
	if x != nil {
		return x.LockingScript
	}
	return nil
}

type LockingScript struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKeyHash []byte `protobuf:"bytes,1,opt,name=pubKeyHash,proto3" json:"pubKeyHash,omitempty"`
}

func (x *LockingScript) Reset() {
	*x = LockingScript{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockingScript) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockingScript) ProtoMessage() {}

func (x *LockingScript) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockingScript.ProtoReflect.Descriptor instead.
func (*LockingScript) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{12}
}

func (x *LockingScript) GetPubKeyHash() []byte {
	if x != nil {
		return x.PubKeyHash
	}
	return nil
}

var File_proto_blockchain_proto protoreflect.FileDescriptor

var file_proto_blockchain_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x44,
	0x75, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x31, 0x0a, 0x09, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x55, 0x0a, 0x09,
	0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x65, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0f, 0x70, 0x65, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64,
	0x64, 0x72, 0x73, 0x22, 0x05, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x22, 0x4d, 0x0a, 0x07, 0x50, 0x72,
	0x65, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69,
	0x67, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x22, 0x4f, 0x0a, 0x09, 0x50, 0x72, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73,
	0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x22, 0x84, 0x01, 0x0a, 0x05, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62,
	0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65,
	0x79, 0x22, 0x76, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x76, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x65, 0x76, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x54, 0x0a, 0x0b, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x54, 0x78, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x07, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x54, 0x78,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x22,
	0x77, 0x0a, 0x07, 0x54, 0x78, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x78, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x74, 0x78, 0x6f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x75, 0x74, 0x78, 0x6f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x3a, 0x0a, 0x0f,
	0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x0f, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x3b, 0x0a, 0x0f, 0x55, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62,
	0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x73, 0x69, 0x67, 0x22, 0x58, 0x0a, 0x08, 0x54, 0x78, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x0d, 0x6c, 0x6f, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22,
	0x2f, 0x0a, 0x0d, 0x4c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x73, 0x68,
	0x32, 0xc9, 0x03, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x0f, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x0a, 0x2e, 0x48,
	0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x1a, 0x0a, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73,
	0x68, 0x61, 0x6b, 0x65, 0x12, 0x27, 0x0a, 0x11, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x1b, 0x0a,
	0x0b, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x06, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x0d, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x50, 0x72,
	0x65, 0x56, 0x6f, 0x74, 0x65, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x23, 0x0a, 0x0f, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x0a,
	0x2e, 0x50, 0x72, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b,
	0x12, 0x30, 0x0a, 0x09, 0x44, 0x75, 0x6d, 0x70, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x75,
	0x6d, 0x70, 0x12, 0x34, 0x0a, 0x0e, 0x44, 0x75, 0x6d, 0x70, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x33, 0x0a, 0x0d, 0x44, 0x75, 0x6d, 0x70,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0a, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x3a, 0x0a,
	0x14, 0x44, 0x75, 0x6d, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x31, 0x0a, 0x0b, 0x44, 0x75, 0x6d,
	0x70, 0x55, 0x54, 0x58, 0x4f, 0x53, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0a, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x44, 0x75, 0x6d, 0x70, 0x42, 0x1e, 0x5a, 0x1c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x68, 0x79, 0x6a,
	0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_blockchain_proto_rawDescOnce sync.Once
	file_proto_blockchain_proto_rawDescData = file_proto_blockchain_proto_rawDesc
)

func file_proto_blockchain_proto_rawDescGZIP() []byte {
	file_proto_blockchain_proto_rawDescOnce.Do(func() {
		file_proto_blockchain_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_blockchain_proto_rawDescData)
	})
	return file_proto_blockchain_proto_rawDescData
}

var file_proto_blockchain_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_blockchain_proto_goTypes = []interface{}{
	(*StringDump)(nil),      // 0: StringDump
	(*BytesDump)(nil),       // 1: BytesDump
	(*Handshake)(nil),       // 2: Handshake
	(*Ack)(nil),             // 3: Ack
	(*PreVote)(nil),         // 4: PreVote
	(*PreCommit)(nil),       // 5: PreCommit
	(*Block)(nil),           // 6: Block
	(*Header)(nil),          // 7: Header
	(*Transaction)(nil),     // 8: Transaction
	(*TxInput)(nil),         // 9: TxInput
	(*UnlockingScript)(nil), // 10: UnlockingScript
	(*TxOutput)(nil),        // 11: TxOutput
	(*LockingScript)(nil),   // 12: LockingScript
	(*empty.Empty)(nil),     // 13: google.protobuf.Empty
}
var file_proto_blockchain_proto_depIdxs = []int32{
	7,  // 0: Block.header:type_name -> Header
	8,  // 1: Block.transactions:type_name -> Transaction
	9,  // 2: Transaction.inputs:type_name -> TxInput
	11, // 3: Transaction.outputs:type_name -> TxOutput
	10, // 4: TxInput.unlockingScript:type_name -> UnlockingScript
	12, // 5: TxOutput.lockingScript:type_name -> LockingScript
	2,  // 6: Node.HandleHandshake:input_type -> Handshake
	8,  // 7: Node.HandleTransaction:input_type -> Transaction
	6,  // 8: Node.HandleBlock:input_type -> Block
	4,  // 9: Node.HandlePreVote:input_type -> PreVote
	5,  // 10: Node.HandlePreCommit:input_type -> PreCommit
	13, // 11: Node.DumpPeers:input_type -> google.protobuf.Empty
	13, // 12: Node.DumpBlockStore:input_type -> google.protobuf.Empty
	13, // 13: Node.DumpBlockList:input_type -> google.protobuf.Empty
	13, // 14: Node.DumpTransactionStore:input_type -> google.protobuf.Empty
	13, // 15: Node.DumpUTXOSet:input_type -> google.protobuf.Empty
	2,  // 16: Node.HandleHandshake:output_type -> Handshake
	3,  // 17: Node.HandleTransaction:output_type -> Ack
	3,  // 18: Node.HandleBlock:output_type -> Ack
	3,  // 19: Node.HandlePreVote:output_type -> Ack
	3,  // 20: Node.HandlePreCommit:output_type -> Ack
	0,  // 21: Node.DumpPeers:output_type -> StringDump
	1,  // 22: Node.DumpBlockStore:output_type -> BytesDump
	1,  // 23: Node.DumpBlockList:output_type -> BytesDump
	1,  // 24: Node.DumpTransactionStore:output_type -> BytesDump
	1,  // 25: Node.DumpUTXOSet:output_type -> BytesDump
	16, // [16:26] is the sub-list for method output_type
	6,  // [6:16] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_blockchain_proto_init() }
func file_proto_blockchain_proto_init() {
	if File_proto_blockchain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_blockchain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringDump); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blockchain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesDump); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blockchain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Handshake); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blockchain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blockchain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreVote); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blockchain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreCommit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blockchain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blockchain_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blockchain_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blockchain_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blockchain_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockingScript); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blockchain_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blockchain_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockingScript); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_blockchain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_blockchain_proto_goTypes,
		DependencyIndexes: file_proto_blockchain_proto_depIdxs,
		MessageInfos:      file_proto_blockchain_proto_msgTypes,
	}.Build()
	File_proto_blockchain_proto = out.File
	file_proto_blockchain_proto_rawDesc = nil
	file_proto_blockchain_proto_goTypes = nil
	file_proto_blockchain_proto_depIdxs = nil
}
