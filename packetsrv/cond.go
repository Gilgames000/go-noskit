package packetsrv

import (
	"github.com/gilgames000/go-noskit/parser"
)

// EntityCondition packet
type EntityCondition struct {
	EntityType        int `json:"entity_type"         parser:"'cond' @String"`
	EntityID          int `json:"entity_id"           parser:"@String"`
	IsAttackAllowed   int `json:"is_attack_allowed"   parser:"@String"`
	IsMovementAllowed int `json:"is_movement_allowed" parser:"@String"`
	Speed             int `json:"speed"               parser:"@String"`
}

// Name of the packet
func (p EntityCondition) Name() string {
	return "cond"
}

// Type of the packet
func (p EntityCondition) Type() parser.PacketType {
	return parser.SERVER
}
