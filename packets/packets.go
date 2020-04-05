package packets

// PacketType represents the type of a packet
type PacketType int

const (
	// CLIENT specifies that a packet has been generated by the client
	CLIENT PacketType = iota
	// SERVER specifies that a packet has been generated by the server
	SERVER
)

// NosPacket is the interface that a packet must implement
type NosPacket interface {
	Name() string
	Type() PacketType
}

// NosPacketStringer is a packet with serialization capabilities
type NosPacketStringer interface {
	NosPacket
	String() string
}
