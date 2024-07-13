package peers

import (
	pb "github.com/fahyjo/blockchain/proto"
)

// Peer represents a peer node in the blockchain network
type Peer struct {
	Height int64         // Height the number of blocks in this Peer's blockchain
	Client pb.NodeClient // client to invoke methods on this Peer's gRPC Server
}

// NewPeer creates a new Peer
func NewPeer(height int64, client pb.NodeClient) *Peer {
	return &Peer{
		Height: height,
		Client: client,
	}
}
