package packetsrv

import "github.com/gilgames000/go-noskit/packets"

// SpawnMob packet
type SpawnMob struct {
	VNum      int `json:"vnum"       parser:"'in' '3' @String"`
	ID        int `json:"id"         parser:" @String"`
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
func (p SpawnMob) Type() packets.PacketType {
	return packets.SERVER
}
