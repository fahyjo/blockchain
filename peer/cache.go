package peer

import (
	"fmt"
	"sync"
)

// PeerCache keeps track of all known peers
type PeerCache struct {
	lock  sync.RWMutex
	Cache map[string]*Peer
}

func NewPeerCache() *PeerCache {
	return &PeerCache{
		lock:  sync.RWMutex{},
		Cache: make(map[string]*Peer),
	}
}

func (c *PeerCache) Get(peerAddr string) (*Peer, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	p, ok := c.Cache[peerAddr]
	if !ok {
		return nil, fmt.Errorf("peer cache miss: %s", peerAddr)
	}
	return p, nil
}

func (c *PeerCache) Put(peerAddr string, p *Peer) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.Cache[peerAddr] = p
}

func (c *PeerCache) Has(peerAddr string) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	_, ok := c.Cache[peerAddr]
	return ok
}

func (c *PeerCache) Size() int {
	return len(c.Cache)
}

func (c *PeerCache) GetPeerAddrs() []string {
	c.lock.RLock()
	defer c.lock.RUnlock()

	peerAddrs := make([]string, 0, c.Size())
	for peerAddr := range c.Cache {
		peerAddrs = append(peerAddrs, peerAddr)
	}
	return peerAddrs
}
