package datastore

import (
	"encoding/binary"
	"io"
)

type RawMapLoader interface {
	Load(mapID int) (io.Reader, error)
}

type MapData struct {
	Width, Height   int
	WalkabilityGrid [][]bool
}

type MapDataStore struct {
	loader RawMapLoader
	maps   map[int]MapData
}

func NewMapDataStore(loader RawMapLoader) MapDataStore {
	return MapDataStore{
		loader: loader,
		maps:   make(map[int]MapData),
	}
}

func (m MapDataStore) loadMapData(mapID int) (MapData, error) {
	r, err := m.loader.Load(mapID)
	if err != nil {
		return MapData{}, err
	}

	var size struct {
		W, H uint16
	}
	if err := binary.Read(r, binary.LittleEndian, &size); err != nil {
		return MapData{}, err
	}

	grid := make([]bool, size.W*size.H)
	if err := binary.Read(r, binary.LittleEndian, &grid); err != nil {
		return MapData{}, err
	}

	var mapData MapData
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

func (m MapDataStore) retrieveMap(mapID int) (MapData, error) {
	if _, ok := m.maps[mapID]; !ok {
		mapData, err := m.loadMapData(mapID)
		if err != nil {
			return MapData{}, err
		}

		m.maps[mapID] = mapData
	}

	return m.maps[mapID], nil
}

func (m MapDataStore) Size(mapID int) (int, int, error) {
	mapData, err := m.retrieveMap(mapID)
	if err != nil {
		return -1, -1, err
	}

	return mapData.Width, mapData.Height, nil
}

func (m MapDataStore) WalkabilityGrid(mapID int) ([][]bool, error) {
	mapData, err := m.retrieveMap(mapID)
	if err != nil {
		return [][]bool{}, err
	}

	return mapData.WalkabilityGrid, nil
}
