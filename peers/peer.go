package peers

import (
	pb "github.com/fahyjo/blockchain/proto"
)

// Peer represents a peer node in the blockchain network
type Peer struct {
	Client pb.NodeClient // client to invoke methods on this Peer's gRPC Server
}

// NewPeer creates a new Peer
func NewPeer(client pb.NodeClient) *Peer {
	return &Peer{
		Client: client,
	}
}
