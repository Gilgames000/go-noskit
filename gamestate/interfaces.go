package gamestate

import (
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/packets"
)

type GameSocket interface {
	Listen(packetNames ...string) <-chan packets.NosPacket
	CloseListener(listener <-chan packets.NosPacket)
	Send(packet string)
}

type Pathfinder interface {
	FindPath(p1 entities.Point, p2 entities.Point, walkabilityGrid [][]bool) ([]entities.Point, error)
	DistanceBetween(p1 entities.Point, p2 entities.Point, walkabilityGrid [][]bool) (int, error)
}

type ItemDataStore interface {
	SearchByVNum(vnum int) (entities.Item, error)
}

type MapDataStore interface {
	Size(mapID int) (int, int, error)
	WalkabilityGrid(mapID int) ([][]bool, error)
}
