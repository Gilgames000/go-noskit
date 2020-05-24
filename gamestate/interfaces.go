package gamestate

import (
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/packets"
)

type PacketSender interface {
	Send(packet ...packets.NosPacketStringer) error
	SendRaw(packet ...string) error
}

type PacketReceiver interface {
	NewListener(packetNames ...string) chan packets.NosPacket
	CloseListener(listener chan packets.NosPacket)
}

// GameSocket provides an abstraction over the low-level implementation
// of the game socket.
type GameSocket interface {
	PacketSender
	PacketReceiver
	Connect(address string, sessionNumber int) error
	Disconnect() error
	IsConnected() bool
}

// Pathfinder provides an abstraction over the implementation of
// pathfinding facilities by requiring a walkability grid as the
// only information needed about the map.
type Pathfinder interface {
	FindPath(p1, p2 entities.Point, walkabilityGrid [][]bool) ([]entities.Point, error)
	DistanceBetween(p1, p2 entities.Point, walkabilityGrid [][]bool) (int, error)
}

type ItemDataStore interface {
	SearchByVNum(vnum int) (entities.Item, error)
}

type MapDataStore interface {
	Size(mapID int) (int, int, error)
	WalkabilityGrid(mapID int) ([][]bool, error)
}
