package gamestate

import (
	"github.com/gilgames000/go-noskit/actions"
	"github.com/gilgames000/go-noskit/entities"
	packetsrv "github.com/gilgames000/go-noskit/packets/server"
)

var _ actions.MapGateway = &MapGateway{}

type MapGateway struct {
	gameSocket   GameSocket
	mapDataStore MapDataStore
	pathfinder   Pathfinder
	mapInfo      entities.Map
}

func NewMapGateway(gameSocket GameSocket, mapDataStore MapDataStore, pathfinder Pathfinder) *MapGateway {
	mapGateway := &MapGateway{
		gameSocket:   gameSocket,
		mapDataStore: mapDataStore,
		pathfinder:   pathfinder,
	}

	go mapGateway.updater()

	return mapGateway
}

func (mg *MapGateway) updater() {
	l := mg.gameSocket.NewListener([]string{
		packetsrv.CharacterPosition{}.Name(),
	}...)
	defer mg.gameSocket.CloseListener(l)

	for {
		packet := <-l
		switch p := packet.(type) {
		case packetsrv.CharacterPosition:
			mg.mapInfo.ID = p.MapID
		default:
			continue
		}
	}
}

func (mg *MapGateway) Info() entities.Map {
	return mg.mapInfo
}

func (mg *MapGateway) DistanceBetween(p1, p2 entities.Point) (int, error) {
	walkabilityGrid, err := mg.mapDataStore.WalkabilityGrid(mg.mapInfo.ID)
	if err != nil {
		return -1, err
	}

	dist, err := mg.pathfinder.DistanceBetween(p1, p2, walkabilityGrid)
	if err != nil {
		return -1, err
	}

	return dist, nil
}

func (mg *MapGateway) FindPath(p1, p2 entities.Point) ([]entities.Point, error) {
	walkabilityGrid, err := mg.mapDataStore.WalkabilityGrid(mg.mapInfo.ID)
	if err != nil {
		return []entities.Point{}, err
	}

	path, err := mg.pathfinder.FindPath(p1, p2, walkabilityGrid)
	if err != nil {
		return []entities.Point{}, err
	}

	return path, nil
}

func (mg *MapGateway) IsWalkable(p entities.Point) bool {
	walkabilityGrid, err := mg.mapDataStore.WalkabilityGrid(mg.mapInfo.ID)
	if err != nil {
		return false
	}

	return walkabilityGrid[p.X][p.Y]
}
