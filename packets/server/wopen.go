package packetsrv

import (
	"github.com/gilgames000/go-noskit/packets"
)

// WindowOpen packet
type WindowOpen struct {
	WindowType int `json:"windows_type" parser:"'wopen' @String"`
	Unknown    int `json:"unknown"      parser:" @String"`
	Unknown2   int `json:"unknown2"     parser:" @String"`
}

// Name of the packet
func (p WindowOpen) Name() string {
	return "wopen"
}

// Type of the packet
func (p WindowOpen) Type() packets.PacketType {
	return packets.SERVER
}
