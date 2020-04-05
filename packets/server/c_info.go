package packetsrv

import (
	"github.com/gilgames000/go-noskit/packets"
)

// CharacterInfo packet
type CharacterInfo struct {
	CharacterName string `json:"character_name" parser:"'c_info' @String"`
	Unknown       string `json:"unknown"        parser:" @String"`
	GroupID       int    `json:"group_id"       parser:" @String"`
	FamilyID      int    `json:"family_id"      parser:" @String"`
	FamilyName    string `json:"family_name"    parser:" @String"`
	CharacterID   int    `json:"character_id"   parser:" @String"`
	Authority     int    `json:"authority"      parser:" @String"`
	Gender        int    `json:"gender"         parser:" @String"`
	Hairstyle     int    `json:"hairstyle"      parser:" @String"`
	HairColor     int    `json:"hair_color"     parser:" @String"`
	Class         int    `json:"class"          parser:" @String"`
	Icon          int    `json:"icon"           parser:" @String"`
	Compliment    int    `json:"compliment"     parser:" @String"`
	Sprite        int    `json:"sprite_id"      parser:" @String"`
	Invisible     int    `json:"invisible"      parser:" @String"`
	FamilyLevel   int    `json:"family_level"   parser:" @String"`
	MorphUpgrade  int    `json:"morph_upgrade"  parser:" @String"`
	ArenaWinner   int    `json:"arena_winner"   parser:" @String"`
	Unknown2      int    `json:"unknown2"       parser:" @String"`
}

// Name of the packet
func (p CharacterInfo) Name() string {
	return "c_info"
}

// Type of the packet
func (p CharacterInfo) Type() packets.PacketType {
	return packets.SERVER
}
