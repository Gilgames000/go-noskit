package actions

import (
	"fmt"
	"github.com/gilgames000/go-noskit/entities"
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

type OutOfRangeError struct {
	msg string
}

func (e OutOfRangeError) Error() string {
	return "character out of range to perform action: " + e.msg
}

type PointNotWalkableError struct {
	p entities.Point
}

func (e PointNotWalkableError) Error() string {
	return fmt.Sprintf("can't walk to the point (%d, %d) for it is not walkable", e.p.X, e.p.Y)
}

func (ci *CharacterInteractor) WalkTo(p entities.Point) error {
	if walkable := ci.Map.IsWalkable(p); !walkable {
		return &PointNotWalkableError{p: p}
	}

	char := ci.Character.Info()
	path := ci.Map.FindPath(char.Position, p)
	if err := ci.Character.Walk(path); err != nil {
		return err
	}

	return nil
}
