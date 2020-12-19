package data

import (
	"fmt"
	"testing"

	"github.com/gilgames000/go-noskit/datastore"
	"github.com/gilgames000/go-noskit/errors"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/afero"
)

var csvToItemDataTests = []struct {
	csvItemsData string
	hasHeader    bool
	itemData     []datastore.ItemData
}{
	{
		`vnum,inventory_pocket
1,0
2,1
3,3`,
		true,
		[]datastore.ItemData{
			{1, 0},
			{2, 1},
			{3, 3},
		},
	},
	{
		`1,0
2,1
3,3`,
		false,
		[]datastore.ItemData{
			{1, 0},
			{2, 1},
			{3, 3},
		},
	},
}

func createItemsFakeFilesystem(csvItemsData string) (afero.Fs, error) {
	fs := afero.NewMemMapFs()

	f, err := fs.Create("items.csv")
	if err != nil {
		return nil, err
	}

	n, err := f.WriteString(csvItemsData)
	if err != nil {
		return nil, err
	} else if n < 1 {
		return nil, errors.New("zero bytes written")
	}

	err = f.Close()

	return fs, err
}

func TestCSVItemsLoad(t *testing.T) {
	for i, tt := range csvToItemDataTests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			fs, err := createItemsFakeFilesystem(tt.csvItemsData)
			if err != nil {
				t.Errorf("%s", err.Error())
				return
			}

			l := NewCSVItemsLoader(fs, "items.csv", tt.hasHeader)
			itemData, err := l.Load()

			if err != nil {
				t.Errorf("test csv to item data failed with error: %s", err.Error())
			} else if !cmp.Equal(itemData, tt.itemData) {
				t.Errorf("test csv to item data failed:\nexpected: %v\nfound: %v\n", tt.itemData, itemData)
			}
		})
	}
}
