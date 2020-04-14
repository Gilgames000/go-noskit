package packetclt

import (
	"fmt"

	"github.com/gilgames000/go-noskit/packets"
)

// NPCRequest packet
type NPCRequest struct {
	EntityType int `json:"entity_type" parser:"'npc_req' @String"`
	EntityID   int `json:"entity_id"   parser:" @String"`
}

// Name of the packet
func (p NPCRequest) Name() string {
	return "npc_req"
}

// Type of the packet
func (p NPCRequest) Type() packets.PacketType {
	return packets.CLIENT
}

// String representation of the packet
func (p NPCRequest) String() string {
	return fmt.Sprintf("%s %d %d", p.Name(), p.EntityType, p.EntityID)
}
