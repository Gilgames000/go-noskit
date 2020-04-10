package packetsrv

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/gilgames000/go-noskit/packets/parser"
	"github.com/google/go-cmp/cmp"
)

func TestShop(t *testing.T) {
	packet := "shop 2 9287 1 0 45 NosBazaar"

	expected := &Shop{
		EntityType: 2,
		EntityID:   9287,
		IsOpen:     1,
		DialogType: 0,
		ShopType:   45,
		ShopName:   "NosBazaar",
	}

	out := &Shop{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {

		t.Errorf("\nshop packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}
