package packetclt

import (
	"fmt"

	"github.com/gilgames000/go-noskit/parser"
)

// NPCRunAction packet
type NPCRunAction struct {
	ActionID       int `json:"packet_field1" parser:"'n_run' @String"`
	ActionModifier int `json:"packet_field2" parser:" @String"`
	EntityType     int `json:"packet_field3" parser:" @String"`
	EntityID       int `json:"packet_field4" parser:" @String"`
}

// Name of the packet
func (p NPCRunAction) Name() string {
	return "n_run"
}

// Type of the packet
func (p NPCRunAction) Type() parser.PacketType {
	return parser.CLIENT
}

// String representation of the packet
func (p NPCRunAction) String() string {
	return fmt.Sprintf("%s %d %d %d %d", p.Name(), p.ActionID, p.ActionModifier, p.EntityType, p.EntityID)
}
