package gamestate

import (
	"github.com/gilgames000/go-noskit/actions"
	"github.com/gilgames000/go-noskit/entities"
)

var _ actions.ItemGateway = &ItemGateway{}

type ItemGateway struct {
	itemDataStore ItemDataStore
}

func (ig *ItemGateway) SearchByVNum(vnum int) (entities.Item, error) {
	return ig.itemDataStore.SearchByVNum(vnum)
}
