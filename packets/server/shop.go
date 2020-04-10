package packetsrv

import (
	"github.com/gilgames000/go-noskit/packets"
)

// Shop packet
type Shop struct {
	EntityType int    `json:"entity_type" parser:"'shop' @String"`
	EntityID   int    `json:"entity_id"   parser:"@String"`
	IsOpen     int    `json:"is_open"     parser:"@String"`
	DialogType int    `json:"dialog_type" parser:"@String"`
	ShopType   int    `json:"shop_type"   parser:"@String?"`
	ShopName   string `json:"shop_name"   parser:"@String?"`
}

// Name of the packet
func (p Shop) Name() string {
	return "shop"
}

// Type of the packet
func (p Shop) Type() packets.PacketType {
	return packets.SERVER
}
