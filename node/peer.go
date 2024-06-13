package node

import (
	pb "github.com/fahyjo/blockchain/proto"
)

type Peer struct {
	Height int32
	Client pb.NodeClient
}

func NewPeer(height int32, client pb.NodeClient) *Peer {
	return &Peer{
		Height: height,
		Client: client,
	}
}
