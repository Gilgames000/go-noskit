package testdoubles

import (
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/errors"
)

type FakeItemGateway struct {
}

func (i FakeItemGateway) SearchByVNum(vnum int) (entities.Item, error) {
	if vnum != 5 {
		return entities.Item{}, errors.ItemNotFound{VNum: vnum}
	}

	return entities.Item{VNum: 5, InventoryPocket: 1}, nil
}

