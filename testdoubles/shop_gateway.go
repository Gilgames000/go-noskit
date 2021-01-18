package testdoubles

import (
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
)

type FakeShopGateway struct {
	BazaarShop entities.Shop
}

func (s FakeShopGateway) All() []entities.Shop {
	return []entities.Shop{s.BazaarShop}
}

func (s FakeShopGateway) SearchByID(id int) (entities.Shop, bool) {
	if id != s.BazaarShop.OwnerID {
		return entities.Shop{}, false
	}

	return s.BazaarShop, true
}

func (s FakeShopGateway) SearchByShopType(st enums.ShopType) (entities.Shop, bool) {
	if st != s.BazaarShop.ShopType {
		return entities.Shop{}, false
	}

	return s.BazaarShop, true
}

