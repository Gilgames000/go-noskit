package entities

import "github.com/gilgames000/go-noskit/enums"

// Shop represents the instance of a shop opened by an NPC or a Player.
type Shop struct {
	OwnerID  int
	Name     string
	ShopType enums.ShopType
}
