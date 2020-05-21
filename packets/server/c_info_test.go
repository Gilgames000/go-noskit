package packetsrv

import (
	"testing"

	"github.com/gilgames000/go-noskit/packets/parser"

	"github.com/alecthomas/participle"
	"github.com/google/go-cmp/cmp"
)

func TestCInfo(t *testing.T) {
	packet := "c_info pyppona - -1 -1 - 1306678 0 1 0 9 0 1 0 0 0 0 0 0 0"

	expected := &CharacterInfo{
		CharacterName: "pyppona",
		Unknown:       "-",
		GroupID:       -1,
		FamilyID:      -1,
		FamilyName:    "-",
		CharacterID:   1306678,
		Authority:     0,
		Gender:        1,
		Hairstyle:     0,
		HairColor:     9,
		Class:         0,
		Icon:          1,
		Compliment:    0,
		Sprite:        0,
		Invisible:     0,
		FamilyLevel:   0,
		MorphUpgrade:  0,
		ArenaWinner:   0,
		Unknown2:      0,
	}

	out := &CharacterInfo{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\nc_info packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}

func TestCInfoError(t *testing.T) {
	packet := "c_info 1 notanumber"

	out := &CharacterInfo{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if err == nil {
		t.Errorf("\nc_info packet parsing should've returned an error\npacket: %+v\nparsed: %+v\nerror: %+v", packet, out, err)
	}
}
