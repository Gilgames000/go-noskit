package data_test

import (
	"fmt"
	"testing"

	"github.com/gilgames000/go-noskit/data"
	"github.com/gilgames000/go-noskit/datastore"
	"github.com/gilgames000/go-noskit/errors"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/afero"
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

func createMapFakeFilesystem(rawMapData []byte) (afero.Fs, error) {
	fs := afero.NewMemMapFs()

	f, err := fs.Create("0")
	if err != nil {
		return nil, err
	}

	n, err := f.Write(rawMapData)
	if err != nil {
		return nil, err
	} else if n < 1 {
		return nil, errors.New("zero bytes written")
	}

	err = f.Close()

	return fs, err
}

func TestRawMapLoad(t *testing.T) {
	for i, tt := range rawToMapDataTests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			fs, err := createMapFakeFilesystem(tt.rawMapData)
			if err != nil {
				t.Errorf("%s", err.Error())
				return
			}

			l := data.NewRawMapLoader(fs, "")

			mapData, err := l.Load(0)
			if err != nil {
				t.Errorf("test raw to map data failed with error: %s", err.Error())
			} else if !cmp.Equal(mapData, tt.mapData) {
				t.Errorf("test raw to map data failed:\nexpected: %v\nfound: %v\n", tt.mapData, mapData)
			}
		})
	}
}
