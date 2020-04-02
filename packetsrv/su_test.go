package packetsrv

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/gilgames000/go-noskit/parser"
	"github.com/google/go-cmp/cmp"
)

func TestSu(t *testing.T) {
	packet := "su 3 1923 1 1306678 0 12 11 200 0 0 1 98 8 0 0"

	expected := &SkillUsage{
		AttackerType:      3,
		AttackerID:        1923,
		TargetType:        1,
		TargetID:          1306678,
		SkillID:           0,
		SkillCooldown:     12,
		AttackAnimation:   11,
		SkillEffect:       200,
		PositionX:         0,
		PositionY:         0,
		TargetIsAlive:     1,
		ResultingHP:       98,
		Damage:            8,
		HitMode:           0,
		SkillTypeMinusOne: 0,
	}

	out := &SkillUsage{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\nsu packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}

func TestSuError(t *testing.T) {
	packet := "su 1 notanumber"

	out := &SkillUsage{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if err == nil {
		t.Errorf("\nsu packet parsing should've returned an error\npacket: %+v\nparsed: %+v\nerror: %+v", packet, out, err)
	}
}
