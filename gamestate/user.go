package gamestate

import (
	"fmt"
	"time"

	"github.com/gilgames000/go-noskit/actions"
	"github.com/gilgames000/go-noskit/errors"
	packetclt "github.com/gilgames000/go-noskit/packets/client"
	packetsrv "github.com/gilgames000/go-noskit/packets/server"
)

var _ actions.UserGateway = &UserGateway{}

type GFClient interface {
	Authenticate(user actions.User, serverLang string) (string, string, error)
	GetLoginCode(user actions.User, token, accountID string) (string, error)
	GetLoginCodeHex(user actions.User, token, accountID string) (string, error)
}

// LoginSocket provides an abstraction over the low-level implementation
// of the login socket.
type LoginSocket interface {
	PacketSender
	PacketReceiver
	Connect(address string) error
	Disconnect() error
	IsConnected() bool
}

type UserGateway struct {
	gfClient       GFClient
	loginSocket    LoginSocket
	gameSocket     GameSocket
	gameClientInfo GameClientGateway
}

// GameClientGateway provides methods to retrieve information about
// the real game client.
type GameClientGateway interface {
	Version() string
	Hash() string
}

func (ug *UserGateway) AuthenticateGFClient(user actions.User, serverLang string) (string, error) {
	token, accountID, err := ug.gfClient.Authenticate(user, serverLang)
	if err != nil {
		return token, err
	}

	return ug.gfClient.GetLoginCodeHex(user, token, accountID)
}

func (ug *UserGateway) ConnectToLoginServer(user actions.User, loginCode, address string, serverNum int) (accountName string, sessionID int, servers []actions.GameServer, err error) {
	err = ug.loginSocket.Connect(address)
	if err != nil {
		return accountName, sessionID, servers, err
	}

	ln := ug.loginSocket.NewListener([]string{
		packetsrv.NsTeST{}.Name(),
		packetsrv.ConnectionFailure{}.Name(),
	}...)
	defer ug.loginSocket.CloseListener(ln)

	err = ug.loginSocket.Send(packetclt.NoS0577{
		LoginCode:        loginCode,
		InstallationUUID: user.InstallationUUID,
		RandomHex:        "0043BA6F", // TODO: randomize
		ServerNumber:     serverNum,
		ClientVersion:    ug.gameClientInfo.Version(),
		ClientHash:       ug.gameClientInfo.Hash(),
	})
	if err != nil {
		return accountName, sessionID, servers, err
	}

	endpoints := make(map[string][]actions.ServerChannel)
	select {
	case res := <-ln:
		switch p := res.(type) {
		case packetsrv.NsTeST:
			accountName = p.Username
			sessionID = p.SessionID
			for _, ep := range p.Endpoints {
				current := endpoints[ep.ServerName]
				endpoints[ep.ServerName] = append(current, actions.ServerChannel{
					Number:  ep.ChannelNumber,
					Address: ep.Address,
				})
			}
		case packetsrv.ConnectionFailure:
			return accountName, sessionID, servers, errors.LoginFailed{
				ErrorCode: fmt.Sprintf("%d", p.Error),
			}
		default:
			return accountName, sessionID, servers, errors.Errorf("server replied with the wrong packet: %s", p.Name())
		}
	case <-time.After(10 * time.Second):
		return accountName, sessionID, servers, errors.ConnectionTimedOut{}
	}

	for k := range endpoints {
		servers = append(servers, actions.GameServer{
			Name:     k,
			Channels: endpoints[k],
		})
	}

	err = ug.loginSocket.Disconnect()

	return accountName, sessionID, servers, err
}

func (ug *UserGateway) ConnectToGameServer(sessionNum, serverNum int, accountName, address string) (characters []actions.AccountCharacter, err error) {
	err = ug.gameSocket.Connect(address, sessionNum)
	if err != nil {
		return characters, err
	}

	ln := ug.gameSocket.NewListener([]string{
		packetsrv.CharacterListItem{}.Name(),
		packetsrv.Fail{}.Name(),
		packetsrv.CharacterListEnd{}.Name(),
	}...)
	defer ug.gameSocket.CloseListener(ln)

	err = ug.gameSocket.SendRaw(fmt.Sprintf("%s GF %d", accountName, serverNum))
	if err != nil {
		return characters, err
	}

	err = ug.gameSocket.SendRaw("thisisgfmode")
	if err != nil {
		return characters, err
	}

	for {
		select {
		case res := <-ln:
			switch p := res.(type) {
			case packetsrv.CharacterListItem:
				characters = append(characters, actions.AccountCharacter{
					Slot: p.Slot,
					Name: p.CharacterName,
				})
			case packetsrv.CharacterListEnd:
				return characters, nil
			case packetsrv.Fail:
				return characters, errors.LoginFailed{ErrorCode: p.Error}
			default:
				return characters, errors.Errorf("server replied with the wrong packet: %s", p.Name())
			}
		case <-time.After(15 * time.Second):
			return characters, errors.ConnectionTimedOut{}
		}
	}
}
