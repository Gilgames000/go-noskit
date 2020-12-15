package datastore

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var maps = map[int][]byte{
	0: {0x03, 0x00, 0x04, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00},
}

type DummyRawMapLoader struct {
	maps map[int][]byte
}

func (l DummyRawMapLoader) Load(mapID int) (io.Reader, error) {
	return bytes.NewReader(l.maps[mapID]), nil
}

var mapDataStoreTests = []struct {
	mapID   int
	mapData MapData
}{
	{
		0,
		MapData{
			Width:  3,
			Height: 4,
			WalkabilityGrid: [][]bool{
				{true, true, true, true},
				{false, true, true, true},
				{true, true, false, true},
			},
		},
	},
}

func TestLoadMapData(t *testing.T) {
	mds := NewMapDataStore(DummyRawMapLoader{maps: maps})
	for i, tt := range mapDataStoreTests {
		t.Run(fmt.Sprintf("walkability test %d", i), func(t *testing.T) {
			data, err := mds.loadMapData(tt.mapID)
			if err != nil {
				t.Errorf("map %d test load map data failed with error: %s", tt.mapID, err.Error())
			} else if !cmp.Equal(data, tt.mapData) {
				t.Errorf("map %d test load map data failed:\nexpected: %v\nfound: %v\n", tt.mapID, tt.mapData, data)
			}
		})
	}
}
