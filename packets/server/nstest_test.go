package packetsrv

import (
	"testing"

	"github.com/gilgames000/go-noskit/packets/parser"

	"github.com/alecthomas/participle"
	"github.com/google/go-cmp/cmp"
)

func TestNsTeST(t *testing.T) {
	packet := "NsTeST  0 user 2 22416 79.110.84.76:4008:0:1.5.UK-NosCitadel 79.110.84.76:4003:2:1.1.UK-NosCitadel -1:-1:-1:10000.10000.1"

	expected := &NsTeST{
		ServerNumber: 0,
		Username:     "user",
		Unknown:      2,
		SessionID:    22416,
		Endpoints: []ServerEndpoint{
			{
				Address:         "79.110.84.76",
				Port:            4008,
				ChannelColor:    0,
				ChannelFullness: 1,
				ChannelNumber:   5,
				ServerName:      "UK-NosCitadel",
			},
			{
				Address:         "79.110.84.76",
				Port:            4003,
				ChannelColor:    2,
				ChannelFullness: 1,
				ChannelNumber:   1,
				ServerName:      "UK-NosCitadel",
			},
			{
				Address:         "-1",
				Port:            -1,
				ChannelColor:    -1,
				ChannelFullness: 10000,
				ChannelNumber:   10000,
				ServerName:      "1",
			},
		},
	}

	out := &NsTeST{}
	err := participle.MustBuild(
		out,
		participle.Lexer(parser.NosLexer()),
		participle.Elide("Whitespace"),
	).ParseString(packet, out)

	if !cmp.Equal(out, expected) || err != nil {
		t.Errorf("\nNsTeST packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet, expected, out, err)
	}
}
