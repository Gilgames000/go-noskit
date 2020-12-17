package data

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gilgames000/go-noskit/datastore"

	"github.com/google/go-cmp/cmp"
)

var csvToItemDataTests = []struct {
	csvItemData string
	hasHeader   bool
	itemData    []datastore.ItemData
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

func TestCSVToItemData(t *testing.T) {
	for i, tt := range csvToItemDataTests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			l := NewCSVItemsLoader("dummy path", tt.hasHeader)
			data, err := l.csvToItemData(bytes.NewReader([]byte(tt.csvItemData)))
			if err != nil {
				t.Errorf("test csv to item data failed with error: %s", err.Error())
			} else if !cmp.Equal(data, tt.itemData) {
				t.Errorf("test csv to item data failed:\nexpected: %v\nfound: %v\n", tt.itemData, data)
			}
		})
	}
}
