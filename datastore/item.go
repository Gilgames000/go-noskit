package datastore

import (
	"strconv"

	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
	"github.com/gilgames000/go-noskit/errors"
)

type CSVItemLoader interface {
	Load() ([][]string, error)
}

type ItemData struct {
	VNum            int
	InventoryPocket enums.InventoryPocket
}

type ItemDataStore struct {
	loader CSVItemLoader
	items  map[int]ItemData
}

func NewItemDataStore(loader CSVItemLoader) (ItemDataStore, error) {
	ds := ItemDataStore{
		loader: loader,
		items:  make(map[int]ItemData),
	}
	err := ds.loadItemData()

	return ds, err
}

func (i ItemDataStore) loadItemData() error {
	csv, err := i.loader.Load()
	if err != nil {
		return err
	}

	for _, line := range csv {
		vnum, err := strconv.Atoi(line[0])
		if err != nil {
			return err
		}

		inventoryPocket, err := strconv.Atoi(line[1])
		if err != nil {
			return err
		}

		i.items[vnum] = ItemData{
			VNum:            vnum,
			InventoryPocket: enums.InventoryPocket(inventoryPocket),
		}
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
