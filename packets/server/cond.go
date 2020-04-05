package packetsrv

import "github.com/gilgames000/go-noskit/packets"

// EntityCondition packet
type EntityCondition struct {
	EntityType int `json:"entity_type" parser:"'cond' @String"`
	EntityID   int `json:"entity_id"   parser:"@String"`
	CanAttack  int `json:"can_attack"  parser:"@String"`
	CanMove    int `json:"can_move"    parser:"@String"`
	Speed      int `json:"speed"       parser:"@String"`
}

// Name of the packet
func (p EntityCondition) Name() string {
	return "cond"
}

// Type of the packet
func (p EntityCondition) Type() packets.PacketType {
	return packets.SERVER
}
