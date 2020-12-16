package data

import (
	"encoding/binary"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gilgames000/go-noskit/datastore"
)

var _ datastore.MapLoader = RawMapLoader{}

type RawMapLoader struct {
	mapsDirectory string
}

func NewRawMapLoader(mapsDirectory string) *RawMapLoader {
	return &RawMapLoader{mapsDirectory: mapsDirectory}
}

func (l RawMapLoader) rawToMapData(r io.Reader) (datastore.MapData, error) {
	var size struct {
		W, H uint16
	}
	if err := binary.Read(r, binary.LittleEndian, &size); err != nil {
		return datastore.MapData{}, err
	}

	grid := make([]bool, size.W*size.H)
	if err := binary.Read(r, binary.LittleEndian, &grid); err != nil {
		return datastore.MapData{}, err
	}

	var mapData datastore.MapData
	mapData.Width = int(size.W)
	mapData.Height = int(size.H)
	mapData.WalkabilityGrid = make([][]bool, mapData.Width)
	for i := range mapData.WalkabilityGrid {
		mapData.WalkabilityGrid[i] = make([]bool, mapData.Height)
	}

	for i := range grid {
		mapData.WalkabilityGrid[i%mapData.Width][i/mapData.Width] = !grid[i]
	}

	return mapData, nil
}

func (l RawMapLoader) Load(mapID int) (datastore.MapData, error) {
	f, err := os.Open(filepath.Join(l.mapsDirectory, strconv.Itoa(mapID)))
	if err != nil {
		return datastore.MapData{}, err
	}

	mapData, err := l.rawToMapData(f)
	if err != nil {
		return datastore.MapData{}, err
	}

	return mapData, nil
}
