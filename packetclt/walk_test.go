package packetclt

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/gilgames000/go-noskit/parser"
	"github.com/google/go-cmp/cmp"
)

func TestWalk(t *testing.T) {
	packet := "walk 17 36 0 11"

	expected := &Walk{
		X:        17,
		Y:        36,
		Checksum: 0,
		Speed:    11,
	}

	out := &Walk{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\nwalk packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}

func TestWalkString(t *testing.T) {
	packet := &Walk{
		X:        17,
		Y:        36,
		Checksum: 0,
		Speed:    11,
	}

	expected := "walk 17 36 0 11"

	out := packet.String()

	if !cmp.Equal(out, expected) {
		t.Errorf("\nwalk packet string representation failed\npacket: %+v\nexpected: %+v\nparsed: %+v", packet, expected, out)
	}
}

func TestWalkError(t *testing.T) {
	packet := "walk 1 notanumber"

	out := &Walk{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if err == nil {
		t.Errorf("\ne_info packet parsing should've returned an error\npacket: %+v\nparsed: %+v\nerror: %+v", packet, out, err)
	}
}
