package packetsrv

import (
	"github.com/gilgames000/go-noskit/parser"
)

// CharacterStatus packet
type CharacterStatus struct {
	CurrentHP int `json:"current_hp" parser:"'stat' @String"`
	MaxHP     int `json:"max_hp"     parser:"@String"`
	CurrentMP int `json:"current_mp" parser:"@String"`
	MaxMP     int `json:"max_mp"     parser:"@String"`
	Unknown   int `json:"unknown"    parser:"@String"`
	Unknown2  int `json:"unknown2"   parser:"@String"`
}

// Name of the packet
func (p CharacterStatus) Name() string {
	return "stat"
}

// Type of the packet
func (p CharacterStatus) Type() parser.PacketType {
	return parser.SERVER
}
