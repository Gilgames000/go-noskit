package packetsrv

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/gilgames000/go-noskit/parser"
	"github.com/google/go-cmp/cmp"
)

func TestSc(t *testing.T) {
	packet := "sc 0 0 52 66 54 4 70 1 0 44 46 65 2 70 0 24 46 28 46 23 0 0 0 0"

	expected := &CharacterEquipment{
		MainWeaponType:       0,
		MainWeaponUp:         0,
		MainWeaponMinDmg:     52,
		MainWeaponMaxDmg:     66,
		MainWeaponHitRate:    54,
		MainWeaponCritRate:   4,
		MainWeaponCritDmg:    70,
		SecondWeaponType:     1,
		SecondWeaponUp:       0,
		SecondWeaponMinDmg:   44,
		SecondWeaponMaxDmg:   46,
		SecondWeaponHitRate:  65,
		SecondWeaponCritRate: 2,
		SecondWeaponCritDmg:  70,
		ArmorUp:              0,
		MeleeDef:             24,
		MeleeDodge:           46,
		RangedDef:            28,
		RangedDodge:          46,
		MagicDef:             23,
		FireRes:              0,
		WaterRes:             0,
		LightRes:             0,
		DarkRes:              0,
	}

	out := &CharacterEquipment{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\nsc packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}

func TestScError(t *testing.T) {
	packet := "sc 1 notanumber"

	out := &CharacterEquipment{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if err == nil {
		t.Errorf("\nsc packet parsing should've returned an error\npacket: %+v\nparsed: %+v\nerror: %+v", packet, out, err)
	}
}
