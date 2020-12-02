package data

import (
	"encoding/csv"
	"os"
	"path/filepath"

	"github.com/gilgames000/go-noskit/datastore"
)

var _ datastore.CSVItemLoader = &CSVItemsLoader{}

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

func (l *CSVItemsLoader) Load() ([][]string, error) {
	f, err := os.Open(filepath.Clean(l.itemsDatPath))
	if err != nil {
		return [][]string{}, err
	}

	r := csv.NewReader(f)
	all, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	if l.hasHeader {
		all = all[1:]
	}

	return all, nil
}
