package packetsrv

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/gilgames000/go-noskit/packets/parser"
	"github.com/google/go-cmp/cmp"
)

func TestEntityCondition(t *testing.T) {
	packet := "cond 1 1879326 0 0 11"

	expected := &EntityCondition{
		EntityType: 1,
		EntityID:   1879326,
		CanAttack:  0,
		CanMove:    0,
		Speed:      11,
	}

	out := &EntityCondition{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {

		t.Errorf("\ncond packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}
