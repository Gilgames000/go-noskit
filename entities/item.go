package entities

import "github.com/gilgames000/go-noskit/enums"

// Item holds the info about a specific item.
type Item struct {
	VNum            int
	InventoryPocket enums.InventoryPocket
}

// ItemInstance represents the instance of an item.
type ItemInstance struct {
	Item
	Amount    int
	OwnerID   int
	OwnerName string
}

// Drop represents the instance of an item on a map.
type Drop struct {
	Thing
	ItemInstance
	IsQuestItem bool
}

// BazaarItem represents the instance of an item listed in the bazaar.
type BazaarItem struct {
	ItemInstance
	Price       int
	MinutesLeft int
	SaleStatus  enums.SaleStatus
	SoldAmount  int
}
