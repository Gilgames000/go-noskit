package testdoubles

import (
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
	"github.com/gilgames000/go-noskit/errors"
)

type FakeBazaarGateway struct {
	isOpen bool
}

func (b *FakeBazaarGateway) Open(int) error {
	b.isOpen = true

	return nil
}

func (b *FakeBazaarGateway) Close() error {
	b.isOpen = false

	return nil
}

func (b FakeBazaarGateway) IsOpen() bool {
	return b.isOpen
}

func (b FakeBazaarGateway) SearchItemsByVNumAndPage(vnums []int, _ int) ([]entities.BazaarItem, error) {
	for i := range vnums {
		if vnums[i] != 5 {
			return nil, errors.ItemNotFound{VNum: vnums[i]}
		}
	}

	return []entities.BazaarItem{{
		ItemInstance: entities.ItemInstance{
			Item: entities.Item{
				VNum:            5,
				InventoryPocket: 1,
			},
			Amount:    1,
			OwnerID:   567,
			OwnerName: "test",
		},
		Price:       100,
		MinutesLeft: 5,
		SaleStatus:  enums.All,
		SoldAmount:  0,
	}}, nil
}

