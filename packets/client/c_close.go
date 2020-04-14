package packetclt

import (
	"fmt"

	"github.com/gilgames000/go-noskit/packets"
)

// CClose packet
type CClose struct {
	Unknown int `json:"unknown" parser:"'c_close' @String"`
}

// Name of the packet
func (p CClose) Name() string {
	return "c_close"
}

// Type of the packet
func (p CClose) Type() packets.PacketType {
	return packets.CLIENT
}

// String representation of the packet
func (p CClose) String() string {
	return fmt.Sprintf("%s %d", p.Name(), p.Unknown)
}
