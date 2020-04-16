package gamestate

import (
	"github.com/gilgames000/go-noskit/actions"
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
	packetsrv "github.com/gilgames000/go-noskit/packets/server"
)

var _ actions.ShopGateway = &ShopGateway{}

type ShopGateway struct {
	gameSocket GameSocket
	shops      map[int]entities.Shop
}

func NewShopGateway(gameSocket GameSocket) *ShopGateway {
	shopGateway := &ShopGateway{
		gameSocket: gameSocket,
		shops:      make(map[int]entities.Shop),
	}

	go shopGateway.updater()

	return shopGateway
}

func (sg *ShopGateway) updater() {
	l := sg.gameSocket.Listen([]string{
		packetsrv.Shop{}.Name(),
	}...)
	defer sg.gameSocket.CloseListener(l)

	for {
		packet := <-l
		switch p := packet.(type) {
		case packetsrv.Shop:
			if p.IsOpen != 0 {
				sg.shops[p.EntityID] = newShop(p)
			} else {
				delete(sg.shops, p.EntityID)
			}
		default:
			continue
		}
	}
}

func (sg *ShopGateway) All() []entities.Shop {
	var shops []entities.Shop

	for _, s := range sg.shops {
		shops = append(shops, s)
	}

	return shops
}

func (sg *ShopGateway) SearchByID(id int) (entities.Shop, bool) {
	s, ok := sg.shops[id]

	return s, ok
}

func (sg *ShopGateway) SearchByShopType(st enums.ShopType) (entities.Shop, bool) {
	for _, s := range sg.shops {
		if s.ShopType == st {
			return s, true
		}
	}

	return entities.Shop{}, false
}

func newShop(shop packetsrv.Shop) entities.Shop {
	return entities.Shop{
		OwnerID:  shop.EntityID,
		Name:     shop.ShopName,
		ShopType: enums.ShopType(shop.ShopType),
	}
}
