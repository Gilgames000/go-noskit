package packetclt

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/gilgames000/go-noskit/parser"
	"github.com/google/go-cmp/cmp"
)

func TestBazaarBuyList(t *testing.T) {
	packet := "c_blist 1 0 0 0 0 0 0 0 4 1 2 3 4"

	expected := &SearchBazaar{
		PageIndex:      1,
		Category:       0,
		SubCategory:    0,
		Level:          0,
		Rarity:         0,
		Upgrade:        0,
		Order:          0,
		Unknown:        0,
		ItemListLength: 4,
		Items:          []int{1, 2, 3, 4},
	}

	out := &SearchBazaar{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\ne_info packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}

func TestBazaarBuyListString(t *testing.T) {
	packet := "c_blist 1 0 0 0 0 0 0 0 4 1 2 3 4"

	out := &SearchBazaar{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out.String(), packet) || err != nil {
		t.Errorf("\nc_blist packet Stringing failed\nexpected: %+v\ngot: %+v", packet, out.String())
	}
}
