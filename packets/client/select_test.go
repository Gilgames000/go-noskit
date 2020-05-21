package packetclt

import (
	"testing"

	"github.com/gilgames000/go-noskit/packets/parser"

	"github.com/alecthomas/participle"
	"github.com/google/go-cmp/cmp"
)

func TestSelectCharacter(t *testing.T) {
	packet := "select 1"
	expected := &SelectCharacter{
		Slot: 1,
	}

	out := &SelectCharacter{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\nselect packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}

func TestSelectCharacterString(t *testing.T) {
	packet := &SelectCharacter{
		Slot: 0,
	}

	expected := "select 0"

	out := packet.String()

	if !cmp.Equal(out, expected) {
		t.Errorf("\nselect packet string representation failed\npacket: %+v\nexpected: %+v\nparsed: %+v", packet, expected, out)
	}
}
