package gamestate

import (
	"github.com/gilgames000/go-noskit/actions"
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
	"time"

	"github.com/gilgames000/go-noskit/errors"
	packetclt "github.com/gilgames000/go-noskit/packets/client"
	packetsrv "github.com/gilgames000/go-noskit/packets/server"
)

var _ actions.BazaarGateway = &BazaarGateway{}

type BazaarGateway struct {
	gameSocket    GameSocket
	itemDataStore ItemDataStore
	isOpen        bool
}

func (bg *BazaarGateway) Open(npcID int) error {
	req := bg.gameSocket.Listen(packetsrv.NPCRequest{}.Name())
	defer bg.gameSocket.CloseListener(req)

	bg.gameSocket.Send(packetclt.NPCRequest{
		EntityType: 2,
		EntityID:   npcID,
	}.String())

	select {
	case <-req:
		break
	case <-time.After(5 * time.Second):
		return &errors.BazaarInteractionError{Msg: "request timeout"}
	}
	time.Sleep(1 * time.Second)

	wop := bg.gameSocket.Listen(packetsrv.WindowOpen{}.Name())
	defer bg.gameSocket.CloseListener(wop)

	bg.gameSocket.Send(packetclt.NPCRunAction{
		ActionID:       60,
		ActionModifier: 0,
		EntityType:     2,
		EntityID:       npcID,
	}.String())

	select {
	case <-wop:
		break
	case <-time.After(5 * time.Second):
		return &errors.BazaarInteractionError{Msg: "request timeout"}
	}
	time.Sleep(1 * time.Second)

	res := bg.gameSocket.Listen(packetsrv.BazaarSearchResults{}.Name())
	defer bg.gameSocket.CloseListener(res)

	bg.gameSocket.Send(packetclt.SearchBazaar{}.String())

	select {
	case <-res:
		break
	case <-time.After(5 * time.Second):
		return &errors.BazaarInteractionError{Msg: "request timeout"}
	}
	time.Sleep(1 * time.Second)

	bg.isOpen = true

	return nil
}

func (bg *BazaarGateway) Close() {
	bg.gameSocket.Send(packetclt.CClose{}.String())
	bg.isOpen = false
}

func (bg *BazaarGateway) IsOpen() bool {
	return bg.isOpen
}

func (bg *BazaarGateway) SearchItemsByVNumAndPage(vnums []int, page int) ([]entities.BazaarItem, error) {
	var items []entities.BazaarItem
	var res packetsrv.BazaarSearchResults

	l := bg.gameSocket.Listen(packetsrv.BazaarSearchResults{}.Name())
	defer bg.gameSocket.CloseListener(l)

	bg.gameSocket.Send(packetclt.SearchBazaar{
		PageIndex:      page,
		ItemListLength: len(vnums),
		Items:          vnums,
	}.String())

	select {
	case p := <-l:
		var ok bool
		res, ok = p.(packetsrv.BazaarSearchResults)
		if !ok {
			return items, errors.BazaarInteractionError{Msg: "wrong packet received: " + p.Name()}
		}
	case <-time.After(5 * time.Second):
		return items, &errors.BazaarInteractionError{Msg: "request timeout"}
	}
	time.Sleep(1 * time.Second)

	for _, it := range res.Items {
		itInfo, err := bg.itemDataStore.SearchByVNum(it.VNum)
		if err != nil {
			continue
		}

		items = append(items, newBazaarItem(it, itInfo.InventoryPocket))
	}

	return items, nil
}

func newBazaarItem(it packetsrv.BazaarItem, pocket enums.InventoryPocket) entities.BazaarItem {
	return entities.BazaarItem{
		ItemInstance: entities.ItemInstance{
			Item: entities.Item{
				VNum:            it.VNum,
				InventoryPocket: pocket,
			},
			Amount:    it.Amount,
			OwnerID:   it.OwnerID,
			OwnerName: it.OwnerName,
		},
		Price:       it.Price,
		MinutesLeft: it.MinutesLeft,
	}
}
