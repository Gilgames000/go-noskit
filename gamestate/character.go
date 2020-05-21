package gamestate

import (
	"math"
	"time"

	"github.com/gilgames000/go-noskit/actions"
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
	"github.com/gilgames000/go-noskit/errors"
	packetclt "github.com/gilgames000/go-noskit/packets/client"
	packetsrv "github.com/gilgames000/go-noskit/packets/server"
)

var _ actions.CharacterGateway = &CharacterGateway{}

type CharacterGateway struct {
	gameSocket   GameSocket
	mapDataStore MapDataStore
	pathfinder   Pathfinder
	character    entities.Character
	mapID        int
	canAttack    bool
	canMove      bool
}

func NewCharacterGateway(gameSocket GameSocket, mapDataStore MapDataStore, pathfinder Pathfinder) *CharacterGateway {
	characterGateway := &CharacterGateway{
		gameSocket:   gameSocket,
		mapDataStore: mapDataStore,
		pathfinder:   pathfinder,
	}

	go characterGateway.updater()

	return characterGateway
}

func (cg *CharacterGateway) updater() {
	l := cg.gameSocket.Listen([]string{
		packetsrv.CharacterInfo{}.Name(),
		packetsrv.CharacterStatus{}.Name(),
		packetsrv.CharacterLevel{}.Name(),
		packetsrv.CharacterPosition{}.Name(),
		packetsrv.EntityCondition{}.Name(),
	}...)
	defer cg.gameSocket.CloseListener(l)

	for {
		packet := <-l
		switch p := packet.(type) {
		case packetsrv.CharacterInfo:
			cg.character.Name = p.CharacterName
			cg.character.ID = p.CharacterID
			cg.character.Class = enums.Class(p.Class)
			cg.character.Gender = enums.Gender(p.Gender)
		case packetsrv.CharacterStatus:
			cg.character.HPPercentage = p.CurrentHP / p.MaxHP * 100
			cg.character.MPPercentage = p.CurrentMP / p.MaxMP * 100
			cg.character.CurrentHP = p.CurrentHP
			cg.character.CurrentMP = p.CurrentMP
			cg.character.MaxHP = p.MaxHP
			cg.character.MaxMP = p.MaxMP
		case packetsrv.CharacterLevel:
			cg.character.CombatLevel = p.CombatLevel
		case packetsrv.CharacterPosition:
			cg.mapID = p.MapID
			cg.character.Position.X = p.X
			cg.character.Position.Y = p.Y
		case packetsrv.EntityCondition:
			if p.CanMove == 0 {
				cg.canMove = true
			} else {
				cg.canMove = false
			}

			if p.CanAttack == 0 {
				cg.canAttack = true
			} else {
				cg.canAttack = false
			}

			cg.character.Speed = p.Speed
		default:
			continue
		}
	}
}

func (cg *CharacterGateway) Info() entities.Character {
	return cg.character
}

func (cg *CharacterGateway) CanMove() bool {
	return cg.canMove
}

func (cg *CharacterGateway) CanAttack() bool {
	return cg.canAttack
}

func (cg *CharacterGateway) WalkTo(p entities.Point) error {
	mapID := cg.mapID
	currPos := cg.character.Position

	walkabilityGrid, err := cg.mapDataStore.WalkabilityGrid(mapID)
	if err != nil {
		return err
	}

	path, err := cg.pathfinder.FindPath(currPos, p, walkabilityGrid)
	if err != nil {
		return err
	}

	magicNumber := 2.55
	for len(path) > 0 {
		speed := cg.character.Speed
		maxWalkDist := int(math.Floor(float64(speed) / magicNumber))

		if walkDist := len(path); walkDist <= maxWalkDist {
			nextPos := path[walkDist-1]
			walkDelay := time.Duration(
				float64(walkDist)*math.Floor(magicNumber/float64(speed)*1000),
			) * time.Millisecond

			return cg.walk(nextPos, walkDelay)
		}

		nextPos := path[maxWalkDist-1]
		path = path[maxWalkDist:]
		walkDelay := time.Duration(
			float64(maxWalkDist)*math.Floor(magicNumber/float64(speed)*1000),
		) * time.Millisecond

		if err = cg.walk(nextPos, walkDelay); err != nil {
			return err
		}
	}

	return nil
}

func (cg *CharacterGateway) walk(p entities.Point, delay time.Duration) error {
	if !cg.canMove {
		return &errors.CharacterCannotMove{}
	}

	err := cg.gameSocket.Send(packetclt.Walk{
		X:        p.X,
		Y:        p.Y,
		Checksum: ((p.X + p.Y) % 3) % 2,
		Speed:    cg.character.Speed,
	})
	if err != nil {
		return err
	}

	time.Sleep(delay)
	cg.character.Position = p

	return nil
}
