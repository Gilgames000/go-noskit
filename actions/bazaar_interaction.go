package actions

import (
	"fmt"
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
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

type BazaarInteractionError struct {
	msg string
}

func (e BazaarInteractionError) Error() string {
	return "error while interacting with the bazaar NPC: " + e.msg
}

type ShopNotFoundError struct {
	st enums.ShopType
}

func (e ShopNotFoundError) Error() string {
	return fmt.Sprintf("there is no Shop with shop type %d on the current map", e.st)
}

type ItemNotFoundError struct {
	vnum int
}

func (e ItemNotFoundError) Error() string {
	return fmt.Sprintf("there is no item with VNum %d", e.vnum)
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
		return &ShopNotFoundError{st: enums.NosBazaar}
	}

	dist := bi.Map.DistanceBetween(char.Position, bazaarShop.Position)
	if dist > 3 {
		return &OutOfRangeError{msg: "the character is too distant from the NosBazaar NPC"}
	}

	if err := bi.Bazaar.Open(); err != nil {
		return &BazaarInteractionError{msg: err.Error()}
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
		return []entities.BazaarItem{}, &BazaarInteractionError{msg: "open the bazaar before searching for items"}
	}

	item, err := bi.ItemRepository.SearchByVNum(vnum)
	if err != nil {
		return []entities.BazaarItem{}, &ItemNotFoundError{vnum: vnum}
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
