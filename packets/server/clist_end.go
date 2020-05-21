package packetsrv

import (
	"github.com/gilgames000/go-noskit/packets"
)

// CharacterListEnd packet
type CharacterListEnd struct {
	CListEnd int `parser:"'clist_end'"`
}

// Name of the packet
func (p CharacterListEnd) Name() string {
	return "clist_end"
}

// Type of the packet
func (p CharacterListEnd) Type() packets.PacketType {
	return packets.SERVER
}
