package packetsrv

import (
	"testing"

	"github.com/gilgames000/go-noskit/packets/parser"

	"github.com/alecthomas/participle"
	"github.com/google/go-cmp/cmp"
)

func TestCharacterListItem(t *testing.T) {
	packet := "clist 1 name 0 1 0 9 0 0 1 0 -1.12.1.8.-1.-1.-1.-1.-1.-1 1  1 1 -1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1.-1. 0 0"

	expected := &CharacterListItem{
		Slot:          1,
		CharacterName: "name",
	}

	out := &CharacterListItem{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\nclist packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}
