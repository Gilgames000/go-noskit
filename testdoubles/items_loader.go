package testdoubles

import "github.com/gilgames000/go-noskit/datastore"

type ItemsLoaderStub struct {
}

func (ItemsLoaderStub) Load() ([]datastore.ItemData, error) {
	return []datastore.ItemData{{VNum: 5, InventoryPocket: 2}}, nil
}

