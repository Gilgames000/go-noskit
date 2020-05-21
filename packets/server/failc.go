package packetsrv

import (
	"github.com/gilgames000/go-noskit/packets"
)

// ConnectionFailure packet
type ConnectionFailure struct {
	Error int `json:"error" parser:"'failc' @String"`
}

// Name of the packet
func (p ConnectionFailure) Name() string {
	return "failc"
}

// Type of the packet
func (p ConnectionFailure) Type() packets.PacketType {
	return packets.SERVER
}
