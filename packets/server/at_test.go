package packetsrv

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/gilgames000/go-noskit/packets/parser"
	"github.com/google/go-cmp/cmp"
)

func TestCharacterPosition(t *testing.T) {
	packet := "at 1317843 147 9 30 2 0 6 1 -1"

	expected := &CharacterPosition{
		CharacterID: 1317843,
		MapID:       147,
		X:           9,
		Y:           30,
		Direction:   2,
		Unknown:     0,
		MusicID:     6,
		Unknown2:    1,
		Unknown3:    -1,
	}

	out := &CharacterPosition{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {

		t.Errorf("\at packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}
