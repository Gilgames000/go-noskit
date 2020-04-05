package packetclt

import (
	"fmt"

	"github.com/gilgames000/go-noskit/packets"
)

// Walk packet
type Walk struct {
	X        int `json:"x"        parser:"'walk' @String"`
	Y        int `json:"y"        parser:" @String"`
	Checksum int `json:"checksum" parser:" @String"`
	Speed    int `json:"speed"    parser:" @String"`
}

// Name of the packet
func (p Walk) Name() string {
	return "walk"
}

// Type of the packet
func (p Walk) Type() packets.PacketType {
	return packets.CLIENT
}

// String representation of the packet
func (p Walk) String() string {
	return fmt.Sprintf("%s %d %d %d %d", p.Name(), p.X, p.Y, p.Checksum, p.Speed)
}
