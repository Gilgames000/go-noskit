package parser

import (
	"reflect"
	"regexp"

	"github.com/gilgames000/go-noskit/errors"
	"github.com/gilgames000/go-noskit/packets"

	"github.com/alecthomas/participle"
)

// NosPacketParser provides a collection of methods to parse nostale packets
// and convert them between different formats or notations. It also holds an
// initially-empty list of parsable packets.
type NosPacketParser struct {
	registeredNamesRegex *regexp.Regexp
	registeredNames      string
	clientPackets        map[string]reflect.Type
	clientParsers        map[string]*participle.Parser
	serverPackets        map[string]reflect.Type
	serverParsers        map[string]*participle.Parser
}

// New initializes and returns a new packet parser.
func New() *NosPacketParser {
	return &NosPacketParser{
		registeredNamesRegex: regexp.MustCompile(""),
		registeredNames:      "",
		clientPackets:        make(map[string]reflect.Type),
		clientParsers:        make(map[string]*participle.Parser),
		serverPackets:        make(map[string]reflect.Type),
		serverParsers:        make(map[string]*participle.Parser),
	}
}

// RegisterPacket adds a new packet to the list of packets that the parser will
// be able to parse.
func (p *NosPacketParser) RegisterPacket(packet packets.NosPacket) error {
	packetStructType := reflect.TypeOf(packet)

	// TODO: if already exists?
	// TODO: race conditions?
	// TODO: unregister?

	switch packet.Type() {
	case packets.CLIENT:
		p.clientPackets[packet.Name()] = packetStructType
	case packets.SERVER:
		p.serverPackets[packet.Name()] = packetStructType
	default:
		return errors.New("invalid packet struct: the 'Type() packets.PacketType' method must return either packets.CLIENT or packets.SERVER")
	}

	// Update the names regex by putting the new name at the beginning
	// if the current regex contains a submatch so that it won't take
	// precedence; put it at the end otherwise
	if p.registeredNamesRegex.Find([]byte(packet.Name())) != nil {
		p.registeredNames = "|" + regexp.QuoteMeta(packet.Name()) + p.registeredNames
	} else {
		p.registeredNames = p.registeredNames + "|" + regexp.QuoteMeta(packet.Name())
	}
	p.registeredNamesRegex = regexp.MustCompile("^(" + p.registeredNames[1:] + ")")

	return nil
}

// ParseClientPacket parses a registered client packet
func (p *NosPacketParser) ParseClientPacket(rawPacket string) (packets.NosPacket, error) {
	return p.parsePacket(rawPacket, packets.CLIENT)
}

// ParseServerPacket parses a registered server packet
func (p *NosPacketParser) ParseServerPacket(rawPacket string) (packets.NosPacket, error) {
	return p.parsePacket(rawPacket, packets.SERVER)
}

func (p *NosPacketParser) parsePacket(rawPacket string, packetType packets.PacketType) (packets.NosPacket, error) {
	var registeredPackets map[string]reflect.Type
	var parsers map[string]*participle.Parser

	switch packetType {
	case packets.CLIENT:
		registeredPackets = p.clientPackets
		parsers = p.clientParsers
	case packets.SERVER:
		registeredPackets = p.serverPackets
		parsers = p.serverParsers
	}

	packetName := string(p.registeredNamesRegex.Find([]byte(rawPacket)))
	packetStructType, ok := registeredPackets[packetName]
	if !ok {
		return nil, errors.Errorf("unknown packet: '%s'", rawPacket)
	}

	packetStructPtr := reflect.New(packetStructType)
	packetParser, ok := parsers[packetName]
	if !ok {
		packetParser = participle.MustBuild(
			packetStructPtr.Interface(),
			participle.Lexer(NosLexer()),
			participle.Elide("Whitespace"),
		)
	}

	err := packetParser.ParseString(rawPacket, packetStructPtr.Interface())
	packetStruct := packetStructPtr.Elem().Interface().(packets.NosPacket)

	return packetStruct, err
}
