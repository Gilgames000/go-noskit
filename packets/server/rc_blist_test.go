package packetsrv

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/gilgames000/go-noskit/packets/parser"
	"github.com/google/go-cmp/cmp"
)

func TestBazaarBuyListResults(t *testing.T) {
	packet := "rc_blist 0 441807|10219|CoolGuy|4903|1|0|10000000|28778|2|0|6|7|0|0|1^4903^6^7^0^92^552^700^476^14^190^86^100^900000^-1^2^10219^2^1.16.33^1.1.77 441807|10219|CoolGuy|4903|1|0|10000000|28778|2|0|6|7|0|0|"

	expected := &BazaarSearchResults{
		0,
		[]BazaarItem{
			{
				441807,
				10219,
				"CoolGuy",
				4903,
				1,
				0,
				10000000,
				28778,
				2,
				0,
				6,
				7,
				0,
				0,
				"1^4903^6^7^0^92^552^700^476^14^190^86^100^900000^-1^2^10219^2^1.16.33^1.1.77",
			},
			{
				441807,
				10219,
				"CoolGuy",
				4903,
				1,
				0,
				10000000,
				28778,
				2,
				0,
				6,
				7,
				0,
				0,
				"",
			},
		},
	}

	out := &BazaarSearchResults{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {

		t.Errorf("\nrc_blist packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}

func TestBazaarEmptyBuyList(t *testing.T) {
	packet := "rc_blist 0"

	expected := &BazaarSearchResults{
		0,
		[]BazaarItem{},
	}

	out := &BazaarSearchResults{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if out.PageIndex != 0 || len(out.Items) != 0 || err != nil {

		t.Errorf("\nrc_blist packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}
