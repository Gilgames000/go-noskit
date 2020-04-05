package packetsrv

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/gilgames000/go-noskit/packets/parser"
	"github.com/google/go-cmp/cmp"
)

func TestCharacterLevel(t *testing.T) {
	packet := "lev 1 0 1 0 300 2200 0 2 0 0 3 0"

	expected := &CharacterLevel{
		CombatLevel:     1,
		CurrentCombatXP: 0,
		JobLevel:        1,
		CurrentJobXP:    0,
		MaxCombatXP:     300,
		MaxJobXP:        2200,
		Reputation:      0,
		SkillCP:         2,
		HeroXP:          0,
		HeroLevel:       0,
		HeroXPLoad:      3,
		Unknown:         0,
	}

	out := &CharacterLevel{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {

		t.Errorf("\nlev packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}
