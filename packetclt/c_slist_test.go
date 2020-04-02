package packetclt

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/gilgames000/go-noskit/parser"
	"github.com/google/go-cmp/cmp"
)

func TestBazaarSellList(t *testing.T) {
	packet := "c_slist 0 0"

	expected := &BazaarUserListings{
		PageIndex: 0,
		Status:    0,
	}

	out := &BazaarUserListings{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\ne_info packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}

func TestBazaarSellListError(t *testing.T) {
	packet := "c_slist 0 0 1"

	out := &BazaarUserListings{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if err == nil {
		t.Errorf("\ne_info packet parsing should've returned an error\npacket: %+v\nparsed: %+v\nerror: %+v", packet, out, err)
	}
}
