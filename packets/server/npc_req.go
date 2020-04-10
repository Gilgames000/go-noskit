package packetsrv

import "github.com/gilgames000/go-noskit/packets"

// NPCRequest packet
type NPCRequest struct {
	EntityType int `json:"entity_type" parser:"'npc_req' @String"`
	EntityID   int `json:"shop_id"     parser:" @String"`
	DialogType int `json:"dialog_id"   parser:" @String"`
}

// Name of the packet
func (p NPCRequest) Name() string {
	return "npc_req"
}

// Type of the packet
func (p NPCRequest) Type() packets.PacketType {
	return packets.SERVER
}
