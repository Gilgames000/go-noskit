package data

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gilgames000/go-noskit/datastore"

	"github.com/google/go-cmp/cmp"
)

var rawToMapDataTests = []struct {
	rawMapData []byte
	mapData    datastore.MapData
}{
	{
		[]byte{0x03, 0x00, 0x04, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00},
		datastore.MapData{
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

func TestRawToMapData(t *testing.T) {
	l := NewRawMapLoader("dummy path")
	for i, tt := range rawToMapDataTests {
		t.Run(fmt.Sprintf("walkability test %d", i), func(t *testing.T) {
			data, err := l.rawToMapData(bytes.NewReader(tt.rawMapData))
			if err != nil {
				t.Errorf("test raw to map data failed with error: %s", err.Error())
			} else if !cmp.Equal(data, tt.mapData) {
				t.Errorf("test raw to map data failed:\nexpected: %v\nfound: %v\n", tt.mapData, data)
			}
		})
	}
}
