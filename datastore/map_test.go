package datastore_test

import (
	"fmt"
	"testing"

	"github.com/gilgames000/go-noskit/datastore"
	"github.com/gilgames000/go-noskit/testdoubles"

	"github.com/google/go-cmp/cmp"
)

var mapDataTests = []struct {
	mapLoader  datastore.MapLoader
	mapID      int
	mapData    datastore.MapData
	shouldWork bool
}{
	{
		testdoubles.MapLoaderStub{},
		2,
		datastore.MapData{
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
		testdoubles.FailingMapLoaderStub{},
		2,
		datastore.MapData{},
		false,
	},
}

func TestMapWalkabilityGrid(t *testing.T) {
	for i, tt := range mapDataTests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ds := datastore.NewMapDataStore(tt.mapLoader)
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
			ds := datastore.NewMapDataStore(tt.mapLoader)
			w, h, err := ds.Size(tt.mapID)
			if tt.shouldWork && err != nil {
				t.Errorf("%s", err.Error())
			} else if tt.shouldWork && (w != tt.mapData.Width || h != tt.mapData.Height) {
				t.Errorf("expected: (%d,%d)\nfound: (%d,%d)\n", tt.mapData.Width, tt.mapData.Height, w, h)
			}
		})
	}
}
