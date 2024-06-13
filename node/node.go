package node

import (
	"context"
	"net"
	"sync"

	proto "github.com/fahyjo/blockchain/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Node struct {
	ListenAddr string
	PeerLock   sync.RWMutex
	Peers      map[string]*Peer

	Height int32

	logger *zap.Logger

	proto.UnimplementedNodeServer
}

func NewNode(listenAddr string, peers map[string]*Peer, height int32, logger *zap.Logger) *Node {
	return &Node{
		ListenAddr: listenAddr,
		Peers:      peers,
		Height:     height,
		logger:     logger,
	}
}

func (n *Node) Start(peerAddrs []string) error {
	go n.peerDiscovery(peerAddrs)

	grpcServer := grpc.NewServer()
	proto.RegisterNodeServer(grpcServer, n)

	lis, err := net.Listen("tcp", n.ListenAddr)
	if err != nil {
		return err
	}

	n.logger.Info("Starting gRPC server", zap.String("listenAddr", n.ListenAddr))
	return grpcServer.Serve(lis)
}

func (n *Node) HandleHandshake(ctx context.Context, peerMsg *proto.Handshake) (*proto.Handshake, error) {
	peerAddr := peerMsg.ListenAddr
	_, ok := n.Peers[peerAddr]
	if !ok && peerAddr != n.ListenAddr {
		client, err := n.dialPeer(peerAddr)
		if err != nil {
			return nil, err
		}
		peer := NewPeer(peerMsg.Height, client)
		n.addPeer(peerAddr, peer)
		go n.peerDiscovery(peerMsg.PeerListenAddrs)
	}

	msg := &proto.Handshake{
		ListenAddr:      n.ListenAddr,
		PeerListenAddrs: n.getPeerAddrs(),
		Height:          n.Height,
	}
	return msg, nil
}

func (n *Node) peerDiscovery(peerAddrs []string) error {
	for _, peerAddr := range peerAddrs {
		_, ok := n.Peers[peerAddr]
		if !ok && peerAddr != n.ListenAddr {
			client, err := n.dialPeer(peerAddr)
			if err != nil {
				return err
			}

			msg := &proto.Handshake{
				ListenAddr:      n.ListenAddr,
				PeerListenAddrs: n.getPeerAddrs(),
				Height:          n.Height,
			}
			peerMsg, err := client.HandleHandshake(context.Background(), msg)
			if err != nil {
				return err
			}

			peer := NewPeer(peerMsg.Height, client)
			n.addPeer(peerAddr, peer)

			err = n.peerDiscovery(peerMsg.PeerListenAddrs)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (n *Node) addPeer(peerAddr string, peer *Peer) {
	n.PeerLock.Lock()
	defer n.PeerLock.Unlock()
	n.logger.Info("Adding peer", zap.String("listenAddr", peerAddr))
	n.Peers[peerAddr] = peer
}

func (n *Node) dialPeer(peerAddr string) (proto.NodeClient, error) {
	conn, err := grpc.NewClient(peerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := proto.NewNodeClient(conn)
	return client, nil
}

func (n *Node) getPeerAddrs() []string {
	var peerAddrs []string
	for peerAddr := range n.Peers {
		peerAddrs = append(peerAddrs, peerAddr)
	}
	return peerAddrs
}
