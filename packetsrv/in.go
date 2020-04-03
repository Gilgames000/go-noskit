package packetsrv

import (
	nospktparser "github.com/gilgames000/go-noskit/parser"
)

// SpawnMob packet
type SpawnMob struct {
	VNum      int `json:"vnum"       parser:"'in' '3' @String"`
	MobID     int `json:"mob_id"     parser:" @String"`
	PositionX int `json:"position_x" parser:" @String"`
	PositionY int `json:"position_y" parser:" @String"`
	Direction int `json:"direction"  parser:" @String"`
	CurrentHP int `json:"current_hp" parser:" @String"`
	CurrentMP int `json:"current_mp" parser:" @String (String)*"` // TODO: add remaining fields if necessary
}

// Name of the packet
func (p SpawnMob) Name() string {
	return "in 3"
}

// Type of the packet
func (p SpawnMob) Type() nospktparser.PacketType {
	return nospktparser.SERVER
}
