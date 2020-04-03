package actions

import (
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
	"github.com/gilgames000/go-noskit/errors"
)

type BazaarInteractor struct {
	Bazaar         BazaarGateway
	Character      CharacterGateway
	Map            MapGateway
	ItemRepository ItemGateway
}

type BazaarGateway interface {
	Open() error
	Close() error
	IsOpen() bool
	SearchItemsByVNum(vnum []entities.Item) []entities.BazaarItem
	SearchItemsByVNumAndPage(vnum []entities.Item, page int) []entities.BazaarItem
	UserListings() []entities.BazaarItem
	UserListingsByStatus(status enums.SaleStatus) []entities.BazaarItem
}

type ItemGateway interface {
	SearchByVNum(vnum int) (entities.Item, error)
}

// Open checks if there is a NosBazaar NPC on the current map. If there's one,
// the character will walk to it and Open the bazaar.
func (bi *BazaarInteractor) Open() error {
	var bazaarShop entities.Shop
	var bazaarFound bool

	char := bi.Character.Info()

	for _, shop := range bi.Map.Shops() {
		if shop.ShopType == enums.NosBazaar {
			bazaarShop = shop
			bazaarFound = true
			break
		}
	}

	if !bazaarFound {
		return &errors.ShopNotFoundError{ShopType: enums.NosBazaar}
	}

	dist := bi.Map.DistanceBetween(char.Position, bazaarShop.Position)
	if dist > 3 {
		return &errors.OutOfRangeError{Msg: "the character is too distant from the NosBazaar NPC"}
	}

	if err := bi.Bazaar.Open(); err != nil {
		return &errors.BazaarInteractionError{Msg: err.Error()}
	}

	return nil
}

func (bi *BazaarInteractor) Close() error {
	if !bi.Bazaar.IsOpen() {
		return nil
	}

	return bi.Bazaar.Close()
}

func (bi *BazaarInteractor) SearchItemByVNumAndPage(vnum int, page int) ([]entities.BazaarItem, error) {
	if !bi.Bazaar.IsOpen() {
		return []entities.BazaarItem{}, &errors.BazaarInteractionError{Msg: "open the bazaar before searching for items"}
	}

	item, err := bi.ItemRepository.SearchByVNum(vnum)
	if err != nil {
		return []entities.BazaarItem{}, &errors.ItemNotFoundError{VNum: vnum}
	}

	if page < 0 {
		page = 0
	}
	results := bi.Bazaar.SearchItemsByVNumAndPage([]entities.Item{item}, page)

	return results, nil
}

func (bi *BazaarInteractor) SearchItemByVNum(vnum int) ([]entities.BazaarItem, error) {
	return bi.SearchItemByVNumAndPage(vnum, 0)
}
