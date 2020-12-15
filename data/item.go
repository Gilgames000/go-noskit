package data

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gilgames000/go-noskit/datastore"
	"github.com/gilgames000/go-noskit/enums"
)

var _ datastore.ItemsLoader = &CSVItemsLoader{}

type CSVItemsLoader struct {
	itemsDatPath string
	hasHeader    bool
}

func NewCSVItemsLoader(itemsDatPath string, hasHeader bool) *CSVItemsLoader {
	return &CSVItemsLoader{
		itemsDatPath: itemsDatPath,
		hasHeader:    hasHeader,
	}
}

func (l *CSVItemsLoader) Load() ([]datastore.ItemData, error) {
	f, err := os.Open(filepath.Clean(l.itemsDatPath))
	if err != nil {
		return []datastore.ItemData{}, err
	}

	r := csv.NewReader(f)
	all, err := r.ReadAll()
	if err != nil {
		return []datastore.ItemData{}, err
	}

	if l.hasHeader {
		all = all[1:]
	}

	var items []datastore.ItemData
	for _, line := range all {
		vnum, err := strconv.Atoi(line[0])
		if err != nil {
			return []datastore.ItemData{}, err
		}

		inventoryPocket, err := strconv.Atoi(line[1])
		if err != nil {
			return []datastore.ItemData{}, err
		}

		items = append(items, datastore.ItemData{
			VNum:            vnum,
			InventoryPocket: enums.InventoryPocket(inventoryPocket),
		})
	}

	return items, nil
}
