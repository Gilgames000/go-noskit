package actions

import (
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
	"github.com/gilgames000/go-noskit/errors"
)

// BazaarInteractor lets you interact with the NosBazaar.
type BazaarInteractor struct {
	itemRepository ItemGateway
	bazaar         BazaarGateway
	character      CharacterGateway
	currentMap     MapGateway
	npcs           NPCGateway
	shops          ShopGateway
}

// BazaarGateway provides an abstraction over low-level methods used to
// perform actions on the NosBazaar.
type BazaarGateway interface {
	Open(npcID int) error
	Close() error
	IsOpen() bool
	SearchItemsByVNumAndPage(vnums []int, page int) ([]entities.BazaarItem, error)
}

// ItemGateway provides methods that retrive informations about the items
// in the game.
type ItemGateway interface {
	SearchByVNum(vnum int) (entities.Item, error)
}

// Open checks if there is a NosBazaar NPC on the current map. If there's one,
// the character will talk to it and open the bazaar.
// An error is returned if the NosBazaar is not on the current map or if it's
// not in talk range.
func (bi *BazaarInteractor) Open() error {
	bazaarShop, ok := bi.shops.SearchByShopType(enums.NosBazaar)
	if !ok {
		return &errors.ShopNotFound{ShopType: enums.NosBazaar}
	}

	bazaarNPC, ok := bi.npcs.SearchByID(bazaarShop.OwnerID)
	if !ok {
		return &errors.NPCNotFound{NPCID: bazaarShop.OwnerID}
	}

	dist, err := bi.currentMap.DistanceBetween(bi.character.Info().Position, bazaarNPC.Position)
	if err != nil {
		return err
	}
	if dist > 3 {
		return &errors.CharacterOutOfRange{Action: "open NosBazaar NPC"}
	}

	if err := bi.bazaar.Open(bazaarNPC.ID); err != nil {
		return &errors.BazaarInteractionError{Msg: err.Error()}
	}

	return nil
}

// Close will close the bazaar if it's currently open.
func (bi *BazaarInteractor) Close() error {
	if !bi.bazaar.IsOpen() {
		return nil
	}

	return bi.bazaar.Close()
}

// SearchItemByVNumAndPage lets lets you search for an item in the bazaar
// by specifying its vnum and the results page index.
func (bi *BazaarInteractor) SearchItemByVNumAndPage(vnum int, page int) ([]entities.BazaarItem, error) {
	if !bi.bazaar.IsOpen() {
		return []entities.BazaarItem{}, &errors.BazaarInteractionError{Msg: "open the bazaar before searching for items"}
	}

	item, err := bi.itemRepository.SearchByVNum(vnum)
	if err != nil {
		return []entities.BazaarItem{}, &errors.ItemNotFound{VNum: vnum}
	}

	if page < 0 {
		page = 0
	}

	return bi.bazaar.SearchItemsByVNumAndPage([]int{item.VNum}, page)
}

// SearchItemByVNumAndPage lets lets you search for an item in the bazaar
// by specifying its vnum.
func (bi *BazaarInteractor) SearchItemByVNum(vnum int) ([]entities.BazaarItem, error) {
	return bi.SearchItemByVNumAndPage(vnum, 0)
}
