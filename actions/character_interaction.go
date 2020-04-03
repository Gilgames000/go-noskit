package actions

import (
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/errors"
)

type CharacterInteractor struct {
	Character CharacterGateway
	Map       MapGateway
}

type CharacterGateway interface {
	Info() entities.Character
	Walk(path []entities.Point) error
}

type MapGateway interface {
	Info() entities.Map
	NPCs() []entities.NPC
	Mobs() []entities.Mob
	Players() []entities.Player
	Shops() []entities.Shop
	DistanceBetween(p1 entities.Point, p2 entities.Point) int
	FindPath(p1 entities.Point, p2 entities.Point) []entities.Point
	IsWalkable(p entities.Point) bool
}

func (ci *CharacterInteractor) WalkTo(p entities.Point) error {
	if walkable := ci.Map.IsWalkable(p); !walkable {
		return &errors.PointNotWalkableError{Point: p}
	}

	char := ci.Character.Info()
	path := ci.Map.FindPath(char.Position, p)
	if err := ci.Character.Walk(path); err != nil {
		return err
	}

	return nil
}
