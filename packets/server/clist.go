package packetsrv

import (
	"github.com/gilgames000/go-noskit/packets"
)

// CharacterListItem packet
type CharacterListItem struct {
	Slot          int    `json:"slot"           parser:"'clist' @String"`
	CharacterName string `json:"character_name" parser:"@String String*"` // TODO: add remaining fields if necessary
}

// Name of the packet
func (p CharacterListItem) Name() string {
	return "clist"
}

// Type of the packet
func (p CharacterListItem) Type() packets.PacketType {
	return packets.SERVER
}
