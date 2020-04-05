package packetsrv

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gilgames000/go-noskit/packets"

	"github.com/alecthomas/participle/lexer"
)

// BazaarSearchResults packet
type BazaarSearchResults struct {
	PageIndex int          `json:"page_index" parser:"'rc_blist' @String"`
	Items     []BazaarItem `json:"items"      parser:"@@*"`
}

// Name of the packet
func (p BazaarSearchResults) Name() string {
	return "rc_blist"
}

// Type of the packet
func (p BazaarSearchResults) Type() packets.PacketType {
	return packets.SERVER
}

// BazaarItem represents an element of the buy list
type BazaarItem struct {
	ListingID   int    `json:"listing_id"`
	OwnerID     int    `json:"owner_id"`
	OwnerName   string `json:"owner_name"`
	ItemID      int    `json:"item_id"`
	Amount      int    `json:"amount"`
	IsBundle    int    `json:"is_bundle"`
	Price       int    `json:"price"`
	MinutesLeft int    `json:"minutes_left"`
	Unknown     int    `json:"unknown"`
	Unknown2    int    `json:"unknown2"`
	Rarity      int    `json:"rarity"`
	Upgrade     int    `json:"upgrade"`
	Unknown3    int    `json:"unknown3"`
	Unknown4    int    `json:"unknown4"`
	Info        string `json:"info"`
}

func (item *BazaarItem) Parse(lex *lexer.PeekingLexer) error {
	token, err := lex.Next()
	if err != nil {
		return err
	}

	if token.EOF() {
		return errors.New("EOF reached")
	}

	split := strings.Split(token.Value, "|")
	if len(split) != 15 {
		return errors.New("syntax error in pipe list: wrong number of elements")
	}

	item.OwnerName = split[2]
	item.Info = split[14]

	iter := []struct {
		idx   int
		field *int
	}{
		{0, &item.ListingID},
		{1, &item.OwnerID},
		{3, &item.ItemID},
		{4, &item.Amount},
		{5, &item.IsBundle},
		{6, &item.Price},
		{7, &item.MinutesLeft},
		{8, &item.Unknown},
		{9, &item.Unknown2},
		{10, &item.Rarity},
		{11, &item.Upgrade},
		{12, &item.Unknown3},
		{13, &item.Unknown4},
	}

	for i := range iter {
		var num int

		if num, err = strconv.Atoi(split[iter[i].idx]); err != nil {
			return err
		}
		*iter[i].field = num
	}

	return nil
}
