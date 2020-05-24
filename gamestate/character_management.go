package gamestate

import (
	"fmt"
	"time"

	"github.com/gilgames000/go-noskit/actions"
	"github.com/gilgames000/go-noskit/enums"
	"github.com/gilgames000/go-noskit/errors"
	packetclt "github.com/gilgames000/go-noskit/packets/client"
	packetsrv "github.com/gilgames000/go-noskit/packets/server"
)

var _ actions.CharacterManagementGateway = &CharacterManagementGateway{}

type CharacterManagementGateway struct {
	gameSocket GameSocket
	status     enums.CharacterStatus
}

func (cmg *CharacterManagementGateway) JoinGame(slot int) error {
	ln := cmg.gameSocket.NewListener([]string{
		packetsrv.OK{}.Name(),
		packetsrv.Info{}.Name(),
	}...)
	defer cmg.gameSocket.CloseListener(ln)

	err := cmg.gameSocket.SendRaw([]string{
		"c_close 0",
		"f_stash_end",
		"c_close 1",
	}...)
	if err != nil {
		return err
	}

	err = cmg.gameSocket.Send(packetclt.SelectCharacter{Slot: slot})
	if err != nil {
		return err
	}

	select {
	case res := <-ln:
		switch p := res.(type) {
		case packetsrv.OK:
			break
		case packetsrv.Info:
			return errors.JoinGameFailed{Slot: slot, Message: p.Message}
		default:
			return errors.Errorf("server replied with the wrong packet: %s", p.Name())
		}
	case <-time.After(15 * time.Second):
		return errors.ConnectionTimedOut{}
	}

	err = cmg.gameSocket.SendRaw([]string{
		"game_start",
		"lbs 0",
		"c_close 1",
		"glist 0 0",
	}...)
	if err != nil {
		return err
	}

	go func() {
		ticker := time.NewTicker(60 * time.Second)
		aliveTime := 0
		for {
			select {
			case <-ticker.C:
				aliveTime += 60
				err = cmg.gameSocket.SendRaw(fmt.Sprintf("pulse %d 0", aliveTime))
				if err != nil {
					return
				}
			}
		}
	}()

	cmg.status = enums.InGame

	return nil
}

func (cmg *CharacterManagementGateway) CurrentStatus() enums.CharacterStatus {
	return cmg.status
}
