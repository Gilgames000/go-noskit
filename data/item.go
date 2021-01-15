package data

import (
	"encoding/csv"
	"io"
	"path/filepath"
	"strconv"

	"github.com/gilgames000/go-noskit/datastore"
	"github.com/gilgames000/go-noskit/enums"

	"github.com/spf13/afero"
)

var _ datastore.ItemsLoader = &CSVItemsLoader{}

type CSVItemsLoader struct {
	filesystem   afero.Fs
	itemsDatPath string
	hasHeader    bool
}

func NewCSVItemsLoader(filesystem afero.Fs, itemsDatPath string, hasHeader bool) *CSVItemsLoader {
	return &CSVItemsLoader{
		filesystem:   filesystem,
		itemsDatPath: itemsDatPath,
		hasHeader:    hasHeader,
	}
}

func CSVToItemData(f io.Reader, hasHeader bool) ([]datastore.ItemData, error) {
	r := csv.NewReader(f)
	all, err := r.ReadAll()
	if err != nil {
		return []datastore.ItemData{}, err
	}

	if hasHeader {
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

func (l *CSVItemsLoader) Load() ([]datastore.ItemData, error) {
	f, err := l.filesystem.Open(filepath.Clean(l.itemsDatPath))
	if err != nil {
		return []datastore.ItemData{}, err
	}

	return CSVToItemData(f, l.hasHeader)
}
