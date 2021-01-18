package datastore_test

import (
	"fmt"
	"testing"

	"github.com/gilgames000/go-noskit/datastore"
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/testdoubles"

	"github.com/google/go-cmp/cmp"
)

var itemSearchByVNumTests = []struct {
	vnum       int
	item       entities.Item
	shouldWork bool
}{
	{
		5,
		entities.Item{
			VNum:            5,
			InventoryPocket: 2,
		},
		true,
	},
	{
		3,
		entities.Item{},
		false,
	},
}

func TestItemSearchByVNum(t *testing.T) {
	for i, tt := range itemSearchByVNumTests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ds, err := datastore.NewItemDataStore(testdoubles.ItemsLoaderStub{})
			if err != nil {
				t.Errorf("%s", err.Error())
				return
			}

			item, err := ds.SearchByVNum(tt.vnum)
			if tt.shouldWork && err != nil {
				t.Errorf("%s", err.Error())
			} else if !cmp.Equal(item, tt.item) {
				t.Errorf("expected: %v\nfound: %v\n", tt.item, item)
			}
		})
	}
}
