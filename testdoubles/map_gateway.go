package testdoubles

import (
	"math"

	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/errors"
)

type FakeMapGateway struct {
}

func (m FakeMapGateway) Info() entities.Map {
	return entities.Map{
		ID:     0,
		Width:  math.MaxInt32,
		Height: math.MaxInt32,
	}
}

func (m FakeMapGateway) DistanceBetween(p1 entities.Point, p2 entities.Point) (int, error) {
	if p1.X == p2.X && p1.Y == p2.Y {
		return 0, nil
	} else if p1.X < 0 || p1.Y < 0 || p2.X < 0 || p2.Y < 0 {
		return -1, errors.NoPathToPoint{From: p1, To: p2}
	}

	return math.MaxInt32, nil
}

func (m FakeMapGateway) FindPath(entities.Point, entities.Point) ([]entities.Point, error) {
	return []entities.Point{}, nil
}

func (m FakeMapGateway) IsWalkable(entities.Point) bool {
	return true
}
