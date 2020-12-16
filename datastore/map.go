package datastore

type MapLoader interface {
	Load(mapID int) (MapData, error)
}

type MapData struct {
	Width, Height   int
	WalkabilityGrid [][]bool
}

type MapDataStore struct {
	loader MapLoader
	maps   map[int]MapData
}

func NewMapDataStore(loader MapLoader) MapDataStore {
	return MapDataStore{
		loader: loader,
		maps:   make(map[int]MapData),
	}
}

func (m MapDataStore) retrieveMap(mapID int) (MapData, error) {
	if _, ok := m.maps[mapID]; !ok {
		mapData, err := m.loader.Load(mapID)
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
