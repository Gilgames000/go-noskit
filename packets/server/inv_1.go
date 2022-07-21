package packetsrv

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gilgames000/go-noskit/packets"

	"github.com/alecthomas/participle/lexer"
)

// BazaarBuyListResults packet
type InvPackageResults1 struct {
	PageIndex int               `json:"page_index" parser:"'inv' @String"`
	Items     []InvPackageItem1 `json:"items"      parser:"@@*"`
}

func (p InvPackageResults1) Type() packets.PacketType {
	return packets.SERVER
}

// Name of the packet
func (p InvPackageResults1) Name() string {
	return "inv 1"
}

// BazaarBuyListItem represents an element of the buy list
type InvPackageItem1 struct {
	Slot   int `json:"slot"`
	ItemID int `json:"item_id"`
	Amount int `json:"amount"`
}

func (p InvPackageItem1) String() string {
	str := fmt.Sprintf(
		"%d %d %d",
		p.Slot,
		p.ItemID,
		p.Amount,
	)

	return fmt.Sprintf("%s", str)
}

func (p InvPackageResults1) String() string {
	str := fmt.Sprintf(
		"%d",
		p.PageIndex,
	)
	list := strings.Trim(fmt.Sprintf("%v", p.Items), "[]")

	return fmt.Sprintf("%s %s", str, list)
}

func (item *InvPackageItem1) Parse(lex *lexer.PeekingLexer) error {
	token, err := lex.Next()
	if err != nil {
		return err
	}

	if token.EOF() {
		return errors.New("EOF reached")
	}

	split := strings.Split(token.Value, ".")
	if len(split) != 3 {
		return errors.New("syntax error in pipe list: wrong number of elements")
	}

	iter := []struct {
		idx   int
		field *int
	}{
		{0, &item.Slot},
		{1, &item.ItemID},
		{2, &item.Amount},
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
