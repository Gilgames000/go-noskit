package packetsrv

import (
	"github.com/gilgames000/go-noskit/packets"
)

// Fail packet
type Fail struct {
	Error string `json:"error" parser:"'fail' @String"`
}

// Name of the packet
func (p Fail) Name() string {
	return "fail"
}

// Type of the packet
func (p Fail) Type() packets.PacketType {
	return packets.SERVER
}
