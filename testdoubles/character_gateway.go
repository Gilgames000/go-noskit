package testdoubles

import "github.com/gilgames000/go-noskit/entities"

type FakeCharacterGateway struct {
	Character entities.Character
}

func (c FakeCharacterGateway) Info() entities.Character {
	return c.Character
}

func (c FakeCharacterGateway) WalkTo(entities.Point) error {
	return nil
}

func (c FakeCharacterGateway) CanMove() bool {
	return true
}

func (c FakeCharacterGateway) CanAttack() bool {
	return true
}

