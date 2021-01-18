package testdoubles

import (
	"github.com/gilgames000/go-noskit/datastore"
	"github.com/gilgames000/go-noskit/errors"
)

type MapLoaderStub struct {
}

func (MapLoaderStub) Load(mapID int) (datastore.MapData, error) {
	return datastore.MapData{
		Width:  2,
		Height: 3,
		WalkabilityGrid: [][]bool{
			{false, true, false},
			{true, true, false},
		},
	}, nil
}

type FailingMapLoaderStub struct {
}

func (FailingMapLoaderStub) Load(mapID int) (datastore.MapData, error) {
	return datastore.MapData{}, errors.New("map not found")
}

