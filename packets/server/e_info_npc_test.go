package packetsrv

import (
	"testing"

	"github.com/gilgames000/go-noskit/packets/parser"

	"github.com/alecthomas/participle"
	"github.com/google/go-cmp/cmp"
)

func TestEInfoNPC(t *testing.T) {
	packet := "e_info 10 25 2 0 0 0 0 36 46 32 4 70 0 20 36 20 36 15 0 0 0 0 175 15 -1 Elite^Bacoom"

	expected := &NPCInfo{
		VNum:           25,
		Level:          2,
		AttributeType:  0,
		DamageType:     0,
		AttributeLevel: 0,
		AttackLevel:    0,
		MinDamage:      36,
		MaxDamage:      46,
		HitRate:        32,
		CriticalDamage: 4,
		CriticalRate:   70,
		DefenseLevel:   0,
		MeleeDef:       20,
		MeleeDodge:     36,
		RangedDef:      20,
		RangedDodge:    36,
		MagicDef:       15,
		FireRes:        0,
		WaterRes:       0,
		LightRes:       0,
		DarkRes:        0,
		MaxHP:          175,
		MaxMP:          15,
		NPCName:        "Elite^Bacoom",
	}

	out := &NPCInfo{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\ne_info packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}

func TestEInfoNPCError(t *testing.T) {
	packet := "e_info 10 notanumber"

	out := &NPCInfo{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if err == nil {
		t.Errorf("\ne_info packet parsing should've returned an error\npacket: %+v\nparsed: %+v\nerror: %+v", packet, out, err)
	}
}
