package packetsrv

import (
	"testing"

	"github.com/gilgames000/go-noskit/packets/parser"

	"github.com/alecthomas/participle"
	"github.com/google/go-cmp/cmp"
)

func TestConnectionFailure(t *testing.T) {
	packet := "failc 3"

	expected := &ConnectionFailure{
		Error: 3,
	}

	out := &ConnectionFailure{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\nfailc packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}
