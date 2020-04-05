package packetclt

import (
	"github.com/gilgames000/go-noskit/packets"
)

// CurrentTarget packet
type CurrentTarget struct {
	EntityType int `json:"entity_type" parser:"'ncif' @String"`
	TargetID   int `json:"target_id"   parser:" @String"`
}

// Name of the packet
func (p CurrentTarget) Name() string {
	return "ncif"
}

// Type of the packet
func (p CurrentTarget) Type() packets.PacketType {
	return packets.CLIENT
}
