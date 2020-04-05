package packetsrv

import "github.com/gilgames000/go-noskit/packets"

// CharacterPosition packet
type CharacterPosition struct {
	CharacterID int `json:"character_id" parser:"'at' @String"`
	MapID       int `json:"map_id"       parser:" @String"`
	X           int `json:"x"            parser:" @String"`
	Y           int `json:"y"            parser:" @String"`
	Direction   int `json:"direction"    parser:" @String"`
	Unknown     int `json:"unknown"      parser:" @String"`
	MusicID     int `json:"music_id"     parser:" @String"`
	Unknown2    int `json:"unknown2"     parser:" @String"`
	Unknown3    int `json:"unknown3"     parser:" @String"`
}

// Name of the packet
func (p CharacterPosition) Name() string {
	return "at"
}

// Type of the packet
func (p CharacterPosition) Type() packets.PacketType {
	return packets.SERVER
}
