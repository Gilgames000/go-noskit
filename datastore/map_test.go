package datastore

import (
	"fmt"
	"testing"

	"github.com/gilgames000/go-noskit/errors"

	"github.com/google/go-cmp/cmp"
)

type MapLoaderStub struct {
}

func (MapLoaderStub) Load(mapID int) (MapData, error) {
	return MapData{
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

func (FailingMapLoaderStub) Load(mapID int) (MapData, error) {
	return MapData{}, errors.New("map not found")
}

var mapDataTests = []struct {
	mapLoader  MapLoader
	mapID      int
	mapData    MapData
	shouldWork bool
}{
	{
		MapLoaderStub{},
		2,
		MapData{
			Width:  2,
			Height: 3,
			WalkabilityGrid: [][]bool{
				{false, true, false},
				{true, true, false},
			},
		},
		true,
	},
	{
		FailingMapLoaderStub{},
		2,
		MapData{},
		false,
	},
}

func TestMapWalkabilityGrid(t *testing.T) {
	for i, tt := range mapDataTests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ds := NewMapDataStore(tt.mapLoader)
			walkabilityGrid, err := ds.WalkabilityGrid(tt.mapID)
			if tt.shouldWork && err != nil {
				t.Errorf("%s", err.Error())
			} else if tt.shouldWork && !cmp.Equal(walkabilityGrid, tt.mapData.WalkabilityGrid) {
				t.Errorf("expected: %v\nfound: %v\n", tt.mapData.WalkabilityGrid, walkabilityGrid)
			}
		})
	}
}
func TestMapSize(t *testing.T) {
	for i, tt := range mapDataTests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ds := NewMapDataStore(tt.mapLoader)
			w, h, err := ds.Size(tt.mapID)
			if tt.shouldWork && err != nil {
				t.Errorf("%s", err.Error())
			} else if tt.shouldWork && (w != tt.mapData.Width || h != tt.mapData.Height) {
				t.Errorf("expected: (%d,%d)\nfound: (%d,%d)\n", tt.mapData.Width, tt.mapData.Height, w, h)
			}
		})
	}
}
