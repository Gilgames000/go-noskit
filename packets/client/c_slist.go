package packetclt

import (
	"github.com/gilgames000/go-noskit/packets"
)

// BazaarUserListings packet
type BazaarUserListings struct {
	PageIndex int `json:"page_index" parser:"'c_slist' @String"`
	Status    int `json:"status"     parser:" @String"`
}

// Name of the packet
func (p BazaarUserListings) Name() string {
	return "c_slist"
}

// Type of the packet
func (p BazaarUserListings) Type() packets.PacketType {
	return packets.CLIENT
}
