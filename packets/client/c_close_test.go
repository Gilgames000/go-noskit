package packetclt

import (
	"testing"

	"github.com/gilgames000/go-noskit/packets/parser"

	"github.com/alecthomas/participle"
	"github.com/google/go-cmp/cmp"
)

func TestCClose(t *testing.T) {
	packet := "c_close 0"

	expected := &CClose{
		Unknown: 0,
	}

	out := &CClose{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\nc_close packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}

func TestCCloseString(t *testing.T) {
	packet := &CClose{
		Unknown: 0,
	}

	expected := "c_close 0"

	out := packet.String()

	if !cmp.Equal(out, expected) {
		t.Errorf("\nc_close packet string representation failed\npacket: %+v\nexpected: %+v\nparsed: %+v", packet, expected, out)
	}
}

func TestCCloseError(t *testing.T) {
	packet := "c_close 1 notanumber"

	out := &CClose{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if err == nil {
		t.Errorf("\nc_close packet parsing should've returned an error\npacket: %+v\nparsed: %+v\nerror: %+v", packet, out, err)
	}
}
