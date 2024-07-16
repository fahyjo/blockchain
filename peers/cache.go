package peers

import (
	"fmt"
	"sync"
)

// PeerCache keeps track of all known Peers
type PeerCache struct {
	lock  sync.RWMutex     // lock read/write mutex that ensures PeerCache is concurrent-safe
	cache map[string]*Peer // cache maps network addresses to Peers
}

// NewPeerCache creates a new PeerCache
func NewPeerCache() *PeerCache {
	return &PeerCache{
		lock:  sync.RWMutex{},
		cache: make(map[string]*Peer),
	}
}

// Get retrieves the Peer with the given network address, returns an error if the Peer with the given network address is not found
func (c *PeerCache) Get(peerAddr string) (*Peer, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	p, ok := c.cache[peerAddr]
	if !ok {
		return nil, fmt.Errorf("error getting from peer cache: peer %s not found", peerAddr)
	}
	return p, nil
}

// Put maps the given network address to the given Peer
func (c *PeerCache) Put(peerAddr string, p *Peer) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cache[peerAddr] = p
}

// Has returns true if the Peer with the given network address is in the PeerCache, and false otherwise
func (c *PeerCache) Has(peerAddr string) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	_, ok := c.cache[peerAddr]
	return ok
}

// Size returns the number of known peers
func (c *PeerCache) Size() int {
	return len(c.cache)
}

// GetPeerAddrs returns the network addresses of all known Peers
func (c *PeerCache) GetPeerAddrs() []string {
	c.lock.RLock()
	defer c.lock.RUnlock()

	peerAddrs := make([]string, 0, c.Size())
	for peerAddr := range c.cache {
		peerAddrs = append(peerAddrs, peerAddr)
	}
	return peerAddrs
}

// Cache returns the PeerCache cache
func (c *PeerCache) Cache() map[string]*Peer {
	return c.cache
}
