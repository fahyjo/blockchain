package node

import (
	pb "github.com/fahyjo/blockchain/proto"
)

type Peer struct {
	Height int64
	Client pb.NodeClient
}

func NewPeer(height int64, client pb.NodeClient) *Peer {
	return &Peer{
		Height: height,
		Client: client,
	}
}
