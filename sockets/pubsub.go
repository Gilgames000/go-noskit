package sockets

import (
	"sync"

	"github.com/gilgames000/go-noskit/packets"
)

type NosPacketPubSub struct {
	mtx       sync.RWMutex
	listeners map[string][]chan packets.NosPacket
}

func NewNosPacketPubSub() *NosPacketPubSub {
	return &NosPacketPubSub{listeners: make(map[string][]chan packets.NosPacket)}
}

func (ps *NosPacketPubSub) Subscribe(packetNames ...string) <-chan packets.NosPacket {
	ln := make(chan packets.NosPacket, 100)

	ps.mtx.Lock()
	for _, name := range packetNames {
		ps.listeners[name] = append(ps.listeners[name], ln)
	}
	ps.mtx.Unlock()

	return ln
}

func (ps *NosPacketPubSub) Unsubscribe(subscriber <-chan packets.NosPacket) {
	var idx int

	ps.mtx.Lock()
	for k := range ps.listeners {
		for idx = range ps.listeners[k] {
			if ps.listeners[k][idx] == subscriber {
				ps.listeners[k] = append(
					ps.listeners[k][:idx],
					ps.listeners[k][idx+1:]...,
				)
				break
			}
		}
	}
	ps.mtx.Unlock()
}

func (ps *NosPacketPubSub) Publish(packet packets.NosPacket) {
	ps.mtx.RLock()
	for _, ch := range ps.listeners[packet.Name()] {
		select {
		case ch <- packet:
		default:
		}
	}
	ps.mtx.RUnlock()
}
