package packetclt

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/gilgames000/go-noskit/parser"
	"github.com/google/go-cmp/cmp"
)

func TestNcif(t *testing.T) {
	packet := "ncif 1 12345"
	expected := &CurrentTarget{
		EntityType: 1,
		TargetID:   12345,
	}

	out := &CurrentTarget{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\nncif packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}

func TestNcifError(t *testing.T) {
	packet := "ncif 1 notanumber"

	out := &CurrentTarget{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if err == nil {
		t.Errorf("\nncif packet parsing should've returned an error\npacket: %+v\nparsed: %+v\nerror: %+v", packet, out, err)
	}
}
