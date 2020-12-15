package datastore

import (
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
	"github.com/gilgames000/go-noskit/errors"
)

type ItemsLoader interface {
	Load() ([]ItemData, error)
}

type ItemData struct {
	VNum            int
	InventoryPocket enums.InventoryPocket
}

type ItemDataStore struct {
	loader ItemsLoader
	items  map[int]ItemData
}

func NewItemDataStore(loader ItemsLoader) (ItemDataStore, error) {
	ds := ItemDataStore{
		loader: loader,
		items:  make(map[int]ItemData),
	}
	err := ds.loadItemData()

	return ds, err
}

func (i ItemDataStore) loadItemData() error {
	items, err := i.loader.Load()
	if err != nil {
		return err
	}

	for _, item := range items {
		i.items[item.VNum] = item
	}

	return nil
}

func (i ItemDataStore) SearchByVNum(vnum int) (entities.Item, error) {
	itemData, ok := i.items[vnum]
	if !ok {
		return entities.Item{}, errors.ItemNotFound{VNum: vnum}
	}

	item := entities.Item{
		VNum:            itemData.VNum,
		InventoryPocket: itemData.InventoryPocket,
	}

	return item, nil
}
