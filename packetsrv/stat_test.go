package packetsrv

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/gilgames000/go-noskit/parser"
	"github.com/google/go-cmp/cmp"
)

func TestCharacterStatus(t *testing.T) {
	packet := "stat 221 221 60 60 0 1024"

	expected := &CharacterStatus{
		CurrentHP: 221,
		MaxHP:     221,
		CurrentMP: 60,
		MaxMP:     60,
		Unknown:   0,
		Unknown2:  1024,
	}

	out := &CharacterStatus{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {

		t.Errorf("\nstat packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}
