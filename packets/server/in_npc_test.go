package packetsrv

import (
	"testing"

	"github.com/gilgames000/go-noskit/packets/parser"

	"github.com/alecthomas/participle"
	"github.com/google/go-cmp/cmp"
)

func TestSpawnNPC(t *testing.T) {
	packet := "in 2 793 9287 9 28 1 100 100 460 0 0 -1 1 0 -1 - 2 -1 0 0 0 0 0 0 0 0 0 0"

	expected := &SpawnNPC{
		VNum:       793,
		ID:         9287,
		PositionX:  9,
		PositionY:  28,
		Direction:  1,
		CurrentHP:  100,
		CurrentMP:  100,
		DialogueID: 460,
	}

	out := &SpawnNPC{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\nin packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}

func TestSpawnNPCError(t *testing.T) {
	packet := "in 2 notanumber"

	out := &SpawnNPC{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if err == nil {
		t.Errorf("\nin packet parsing should've returned an error\npacket: %+v\nparsed: %+v\nerror: %+v", packet, out, err)
	}
}
