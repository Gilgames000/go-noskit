package packetsrv

import (
	"testing"

	"github.com/gilgames000/go-noskit/packets/parser"

	"github.com/alecthomas/participle"
	"github.com/google/go-cmp/cmp"
)

func TestNsTeST(t *testing.T) {
	packet := "NsTeST  2 user 2 3 1 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 1 1 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 -99 0 50119 79.110.84.250:4016:1:1.7.Cosmos 79.110.84.250:4015:1:1.6.Cosmos 79.110.84.250:4014:1:1.5.Cosmos 79.110.84.250:4013:1:1.4.Cosmos 79.110.84.250:4012:2:1.3.Cosmos 79.110.84.250:4011:15:1.2.Cosmos 79.110.84.250:4010:4:1.1.Cosmos -1:-1:-1:10000.10000.1"

	expected := &NsTeST{
		ServerNumber: 2,
		Username:     "user",
		Unknown0:     2,
		Unknown1:     3,
		Unknown2:     1,
		WeirdValues:  []int{},
		SessionID:    50119,
		Endpoints: []ServerEndpoint{
			{
				Address:         "79.110.84.250",
				Port:            4016,
				ChannelColor:    1,
				ChannelFullness: 1,
				ChannelNumber:   7,
				ServerName:      "Cosmos",
			},
			{
				Address:         "79.110.84.250",
				Port:            4015,
				ChannelColor:    1,
				ChannelFullness: 1,
				ChannelNumber:   6,
				ServerName:      "Cosmos",
			},
			{
				Address:         "79.110.84.250",
				Port:            4014,
				ChannelColor:    1,
				ChannelFullness: 1,
				ChannelNumber:   5,
				ServerName:      "Cosmos",
			},
			{
				Address:         "79.110.84.250",
				Port:            4013,
				ChannelColor:    1,
				ChannelFullness: 1,
				ChannelNumber:   4,
				ServerName:      "Cosmos",
			},
			{
				Address:         "79.110.84.250",
				Port:            4012,
				ChannelColor:    2,
				ChannelFullness: 1,
				ChannelNumber:   3,
				ServerName:      "Cosmos",
			},
			{
				Address:         "79.110.84.250",
				Port:            4011,
				ChannelColor:    15,
				ChannelFullness: 1,
				ChannelNumber:   2,
				ServerName:      "Cosmos",
			},
			{
				Address:         "79.110.84.250",
				Port:            4010,
				ChannelColor:    4,
				ChannelFullness: 1,
				ChannelNumber:   1,
				ServerName:      "Cosmos",
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
