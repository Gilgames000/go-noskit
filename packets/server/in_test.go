package packetsrv

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/gilgames000/go-noskit/packets/parser"
	"github.com/google/go-cmp/cmp"
)

func TestInMob(t *testing.T) {
	packet := "in 3 24 1878 22 144 2 100 100 0 0 0 -1 1 0 -1 - 2 -1 0 0 0 0 0 0 0 0"

	expected := &SpawnMob{
		VNum:      24,
		MobID:     1878,
		PositionX: 22,
		PositionY: 144,
		Direction: 2,
		CurrentHP: 100,
		CurrentMP: 100,
	}

	out := &SpawnMob{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\nin packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}

func TestInMobError(t *testing.T) {
	packet := "in 3 notanumber"

	out := &SpawnMob{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if err == nil {
		t.Errorf("\nin packet parsing should've returned an error\npacket: %+v\nparsed: %+v\nerror: %+v", packet, out, err)
	}
}
