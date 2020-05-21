package packetsrv

import (
	"testing"

	"github.com/gilgames000/go-noskit/packets/parser"

	"github.com/alecthomas/participle"
	"github.com/google/go-cmp/cmp"
)

func TestInfo(t *testing.T) {
	packet := "info My message"

	expected := &Info{
		Message: "My message",
	}

	out := &Info{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\ninfo packet parsing Infoed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}
