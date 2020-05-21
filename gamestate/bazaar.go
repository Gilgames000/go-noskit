package gamestate

import (
	"time"

	"github.com/gilgames000/go-noskit/actions"
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
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

	err := bg.gameSocket.Send(packetclt.NPCRequest{
		EntityType: 2,
		EntityID:   npcID,
	})
	if err != nil {
		return err
	}

	select {
	case <-req:
		break
	case <-time.After(5 * time.Second):
		return &errors.BazaarInteractionError{Msg: "request timeout"}
	}
	time.Sleep(1 * time.Second)

	wop := bg.gameSocket.Listen(packetsrv.WindowOpen{}.Name())
	defer bg.gameSocket.CloseListener(wop)

	err = bg.gameSocket.Send(packetclt.NPCRunAction{
		ActionID:       60,
		ActionModifier: 0,
		EntityType:     2,
		EntityID:       npcID,
	})
	if err != nil {
		return err
	}

	select {
	case <-wop:
		break
	case <-time.After(5 * time.Second):
		return &errors.BazaarInteractionError{Msg: "request timeout"}
	}
	time.Sleep(1 * time.Second)

	res := bg.gameSocket.Listen(packetsrv.BazaarSearchResults{}.Name())
	defer bg.gameSocket.CloseListener(res)

	err = bg.gameSocket.Send(packetclt.SearchBazaar{})
	if err != nil {
		return err
	}

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

func (bg *BazaarGateway) Close() error {
	err := bg.gameSocket.Send(packetclt.CClose{})
	if err != nil {
		return err
	}
	bg.isOpen = false

	return nil
}

func (bg *BazaarGateway) IsOpen() bool {
	return bg.isOpen
}

func (bg *BazaarGateway) SearchItemsByVNumAndPage(vnums []int, page int) ([]entities.BazaarItem, error) {
	var items []entities.BazaarItem
	var res packetsrv.BazaarSearchResults

	l := bg.gameSocket.Listen(packetsrv.BazaarSearchResults{}.Name())
	defer bg.gameSocket.CloseListener(l)

	err := bg.gameSocket.Send(packetclt.SearchBazaar{
		PageIndex:      page,
		ItemListLength: len(vnums),
		Items:          vnums,
	})
	if err != nil {
		return nil, err
	}

	select {
	case p := <-l:
		var ok bool
		res, ok = p.(packetsrv.BazaarSearchResults)
		if !ok {
			return nil, errors.BazaarInteractionError{Msg: "wrong packet received: " + p.Name()}
		}
	case <-time.After(5 * time.Second):
		return nil, &errors.BazaarInteractionError{Msg: "request timeout"}
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
