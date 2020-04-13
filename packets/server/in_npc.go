package packetsrv

import "github.com/gilgames000/go-noskit/packets"

// SpawnNPC packet
type SpawnNPC struct {
	VNum       int `json:"vnum"        parser:"'in' '2' @String"`
	ID         int `json:"id"          parser:" @String"`
	PositionX  int `json:"position_x"  parser:" @String"`
	PositionY  int `json:"position_y"  parser:" @String"`
	Direction  int `json:"direction"   parser:" @String"`
	CurrentHP  int `json:"current_hp"  parser:" @String"`
	CurrentMP  int `json:"current_mp"  parser:" @String"`
	DialogueID int `json:"dialogue_id" parser:" @String (String)*"` // TODO: add remaining fields if necessary
}

// Name of the packet
func (p SpawnNPC) Name() string {
	return "in 2"
}

// Type of the packet
func (p SpawnNPC) Type() packets.PacketType {
	return packets.SERVER
}
