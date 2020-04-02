package packetclt

import (
	"fmt"
	"strings"

	"github.com/gilgames000/go-noskit/parser"
)

// SearchBazaar packet
type SearchBazaar struct {
	PageIndex      int   `json:"page_index"       parser:"'c_blist' @String"`
	Category       int   `json:"category"         parser:" @String"`
	SubCategory    int   `json:"sub_category"     parser:" @String"`
	Level          int   `json:"level"            parser:" @String"`
	Rarity         int   `json:"rarity"           parser:" @String"`
	Upgrade        int   `json:"upgrade"          parser:" @String"`
	Order          int   `json:"order"            parser:" @String"`
	Unknown        int   `json:"unknown"          parser:" @String"`
	ItemListLength int   `json:"item_list_length" parser:" @String"`
	Items          []int `json:"items"            parser:" (@String)+"`
}

// Name of the packet
func (p SearchBazaar) Name() string {
	return "c_blist"
}

// Type of the packet
func (p SearchBazaar) Type() parser.PacketType {
	return parser.CLIENT
}

// String representation of the packet
func (p SearchBazaar) String() string {
	str := fmt.Sprintf(
		"%s %d %d %d %d %d %d %d %d %d",
		p.Name(),
		p.PageIndex,
		p.Category,
		p.SubCategory,
		p.Level,
		p.Rarity,
		p.Upgrade,
		p.Order,
		p.Unknown,
		p.ItemListLength,
	)
	list := strings.Trim(fmt.Sprintf("%v", p.Items), "[]")

	return fmt.Sprintf("%s %s", str, list)
}
