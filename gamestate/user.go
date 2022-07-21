package gamestate

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gilgames000/go-noskit/actions"
	"github.com/gilgames000/go-noskit/enums"
	"github.com/gilgames000/go-noskit/errors"
	packetclt "github.com/gilgames000/go-noskit/packets/client"
	packetsrv "github.com/gilgames000/go-noskit/packets/server"
	gfclient_auth "github.com/stdLemon/nostale-auth"
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

func NewUserGateway(gfClient GFClient, loginSocket LoginSocket, gameSocket GameSocket, gameClientInfo GameClientGateway) *UserGateway {
	return &UserGateway{
		gfClient:       gfClient,
		loginSocket:    loginSocket,
		gameSocket:     gameSocket,
		gameClientInfo: gameClientInfo,
	}
}

type GfAccountData struct {
	Email    string
	Password string
	Locale   string
	Name     string
}

// GameClientGateway provides methods to retrieve information about
// the real game client.
type GameClientGateway interface {
	Version() string
	Hash() string
}

func (ug *UserGateway) AuthenticateGFClient(jsonAccountPath string, jsonIdentityPath string) (string, error) {
	content, err := ioutil.ReadFile(jsonAccountPath)
	if err != nil {
		fmt.Println(err)
	}

	identity_manager, err := gfclient_auth.NewIdentityManager(jsonIdentityPath)
	if err != nil {
		fmt.Println(err)
	}

	account_data := new(GfAccountData)
	json.Unmarshal(content, account_data)

	identity := identity_manager.Get()

	gfclient := gfclient_auth.NewGfClient(
		identity.Fingerprint.UserAgent,
		"Chrome/C2.2.23.1813 (49c0acbee1)",
		identity.Installation_id,
	)

	err = gfclient.Auth(account_data.Email, account_data.Password, account_data.Locale)
	if err != nil {
		fmt.Println(err)
	}

	// Get game account from user mail
	game_account_list, err := gfclient.GetGameAccounts()
	if err != nil {
		fmt.Println(err)
	}

	// Find game account by its name from account file
	game_account, err := gfclient_auth.FindGameAccount(account_data.Name, game_account_list)
	if err != nil {
		fmt.Println(err)
	}

	err = gfclient.Iovation(identity_manager, game_account.Id)
	if err != nil {
		fmt.Println(err)
	}

	code, err := gfclient.Codes(identity_manager, game_account.Id, game_account.GameId)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("code: " + code)
	identity_manager.Save()

	encoded_code := strings.ToUpper(hex.EncodeToString([]byte(code)))
	// fmt.Println(encoded_code)

	return encoded_code, err
}

func (ug *UserGateway) ConnectToLoginServer(user actions.User, loginCode, address string, countryID enums.CountryID) (accountName string, sessionID int, servers []actions.GameServer, err error) {
	err = ug.loginSocket.Connect(address)
	if err != nil {
		return "", -1, []actions.GameServer{}, err
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
		CountryID:        int(countryID),
		ClientVersion:    ug.gameClientInfo.Version(),
		ClientHash:       ug.gameClientInfo.Hash(),
	})
	if err != nil {
		return "", -1, []actions.GameServer{}, err
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
					Port:    ep.Port,
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

func (ug *UserGateway) ConnectToGameServer(sessionNum int, countryID enums.CountryID, accountName, address string) (characters []actions.AccountCharacter, err error) {
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

	err = ug.gameSocket.SendRaw(fmt.Sprintf("%s GF %d", accountName, countryID))
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
