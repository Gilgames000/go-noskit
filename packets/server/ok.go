package packetsrv

import (
	"github.com/gilgames000/go-noskit/packets"
)

// OK packet
type OK struct {
	Ok int `parser:"'OK'"`
}

// Name of the packet
func (p OK) Name() string {
	return "OK"
}

// Type of the packet
func (p OK) Type() packets.PacketType {
	return packets.SERVER
}
