package packetclt

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/gilgames000/go-noskit/parser"
	"github.com/google/go-cmp/cmp"
)

func TestNRun(t *testing.T) {
	packet := "n_run 16 0 2 1234"

	expected := &NPCRunAction{
		ActionID:       16,
		ActionModifier: 0,
		EntityType:     2,
		EntityID:       1234,
	}

	out := &NPCRunAction{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\ne_info packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}

func TestNRunError(t *testing.T) {
	packet := "n_run 1 notanumber"

	out := &NPCRunAction{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if err == nil {
		t.Errorf("\ne_info packet parsing should've returned an error\npacket: %+v\nparsed: %+v\nerror: %+v", packet, out, err)
	}
}
