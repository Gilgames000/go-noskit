package parser

import (
	"testing"

	"github.com/gilgames000/go-noskit/packets"

	"github.com/google/go-cmp/cmp"
)

type MockPacket1 struct {
	Field1 int `parser:"'mockpacket1' @String"`
	Field2 int `parser:" @String"`
}

func (p MockPacket1) Name() string {
	return "mockpacket1"
}

func (p MockPacket1) Type() packets.PacketType {
	return packets.SERVER
}

type MockPacket2 struct {
	Field1 int `parser:"'mockpacket1' '10' @String"`
	Field2 int `parser:" @String"`
}

func (p MockPacket2) Name() string {
	return "mockpacket1 10"
}

func (p MockPacket2) Type() packets.PacketType {
	return packets.SERVER
}

type MockPacket3 struct {
	Field1 string `parser:"'mockpacket3' @String*"`
}

func (p MockPacket3) Name() string {
	return "mockpacket3"
}

func (p MockPacket3) Type() packets.PacketType {
	return packets.SERVER
}

func TestNosPacketParser(t *testing.T) {
	parser := New()
	packet1 := "mockpacket1 5 20"
	packet2 := "mockpacket1 10 20 30"
	packet3 := "mockpacket3"

	parser.RegisterPacket(MockPacket1{})
	parser.RegisterPacket(MockPacket2{})
	parser.RegisterPacket(MockPacket3{})

	expected1 := MockPacket1{
		Field1: 5,
		Field2: 20,
	}

	expected2 := MockPacket2{
		Field1: 20,
		Field2: 30,
	}

	expected3 := MockPacket3{
		Field1: "",
	}

	out1, err1 := parser.ParseServerPacket(packet1)
	out2, err2 := parser.ParseServerPacket(packet2)
	out3, err3 := parser.ParseServerPacket(packet3)

	if !cmp.Equal(out1, expected1) || err1 != nil {
		t.Errorf("\nmockpacket1 packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet1, expected1, out1, err1)
	}

	if !cmp.Equal(out2, expected2) || err2 != nil {
		t.Errorf("\nmockpacket2 packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet2, expected2, out2, err2)
	}

	if !cmp.Equal(out3, expected3) || err3 != nil {
		t.Errorf("\nmockpacket3 packet parsing failed\npacket: %+v\nexpected: %+v\nparsed: %+v\nerror: %+v", packet3, expected3, out3, err3)
	}
}
