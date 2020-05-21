package packetclt

import (
	"fmt"

	"github.com/gilgames000/go-noskit/packets"
)

// SelectCharacter packet
type SelectCharacter struct {
	Slot int `json:"slot" parser:"'select' @String"`
}

// Name of the packet
func (p SelectCharacter) Name() string {
	return "select"
}

// Type of the packet
func (p SelectCharacter) Type() packets.PacketType {
	return packets.CLIENT
}

// String representation of the packet
func (p SelectCharacter) String() string {
	return fmt.Sprintf("%s %d", p.Name(), p.Slot)
}
