package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	peers         map[net.Addr]Peer
	mu            sync.RWMutex
}
