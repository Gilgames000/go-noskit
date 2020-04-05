package packetsrv

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/gilgames000/go-noskit/packets/parser"
	"github.com/google/go-cmp/cmp"
)

func TestWindowOpen(t *testing.T) {
	packet := "wopen 32 0 0"

	expected := &WindowOpen{
		WindowType: 32,
		Unknown:    0,
		Unknown2:   0,
	}

	out := &WindowOpen{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {

		t.Errorf("\nwopen packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}
