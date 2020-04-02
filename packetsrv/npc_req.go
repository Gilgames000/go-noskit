package packetsrv

import (
	"github.com/gilgames000/go-noskit/parser"
)

// NPCRequest packet
type NPCRequest struct {
	EntityType int `json:"entity_type" parser:"'npc_req' @String"`
	ShopID     int `json:"shop_id"     parser:" @String"`
	DialogID   int `json:"dialog_id"   parser:" @String"`
}

// Name of the packet
func (p NPCRequest) Name() string {
	return "npc_req"
}

// Type of the packet
func (p NPCRequest) Type() parser.PacketType {
	return parser.SERVER
}
