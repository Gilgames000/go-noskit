package actions

import (
	"fmt"

	"github.com/gilgames000/go-noskit/enums"
)

// User holds the information about the current user.
type User struct {
	Email            string
	Username         string
	Password         string
	InstallationUUID string
	Locale           string
}

// UserInteractor lets you interact with the current user.
type UserInteractor struct {
	user          UserGateway
	countryID     enums.CountryID
	accountName   string
	sessionNumber int
}

func NewUserInteractor(userGateway UserGateway) *UserInteractor {
	return &UserInteractor{user: userGateway}
}

// UserGateway provides low-level methods to authenticate and log the user
// in the specified NosTale server.
type UserGateway interface {
	// AuthenticateGFClient authenticates the user simulating the real
	// Gameforge Client and returns a login code.
	AuthenticateGFClient(user User, serverLang string) (string, error)
	// ConnectToLoginServer connects the user to the login server by
	// using the login code retrieved from the GF servers.
	// If the authentication is successful, a list of game servers will be
	// returned alongside a session number and the account name necessary to
	// connect to the game.
	ConnectToLoginServer(user User, loginCode, address string, countryID enums.CountryID) (string, int, []GameServer, error)
	// ConnectToGameServer connects the user to the game server using
	// the session number provided by the login server.
	// If the authentication is successful, a list of characters currently
	// present on the account will be returned.
	ConnectToGameServer(sessionNum int, countryID enums.CountryID, accountName, address string) ([]AccountCharacter, error)
}

// GameServer holds the information about an available server.
type GameServer struct {
	Name     string
	Channels []ServerChannel
}

// ServerChannel holds the information about a channel on a server.
type ServerChannel struct {
	Number  int
	Address string
	Port    int
}

// AccountCharacter holds the information about a character on the logged-in
// user account.
type AccountCharacter struct {
	Slot int
	Name string
}

func (ui *UserInteractor) Login(user User, serverLang string, countryID enums.CountryID) ([]GameServer, error) {
	loginCode, err := ui.user.AuthenticateGFClient(user, serverLang)
	if err != nil {
		return []GameServer{}, err
	}

	accountName, sessionNum, servers, err := ui.user.ConnectToLoginServer(
		user,
		loginCode,
		"login.nostale.gfsrv.net:"+string(getLoginPort(countryID)),
		countryID,
	)
	if err != nil {
		return []GameServer{}, err
	}

	ui.countryID = countryID
	ui.accountName = accountName
	ui.sessionNumber = sessionNum

	return servers, err
}

func (ui *UserInteractor) Connect(channel ServerChannel) ([]AccountCharacter, error) {
	characters, err := ui.user.ConnectToGameServer(
		ui.sessionNumber,
		ui.countryID,
		ui.accountName,
		fmt.Sprintf("%s:%d", channel.Address, channel.Port),
	)
	if err != nil {
		return []AccountCharacter{}, err
	}

	return characters, err
}

func getLoginPort(countryID enums.CountryID) enums.LoginPort {
	switch countryID {
	case enums.EN:
		return enums.LoginPortEN
	case enums.DE:
		return enums.LoginPortDE
	case enums.FR:
		return enums.LoginPortFR
	case enums.IT:
		return enums.LoginPortIT
	case enums.PL:
		return enums.LoginPortPL
	case enums.ES:
		return enums.LoginPortES
	case enums.CZ:
		return enums.LoginPortCZ
	case enums.RU:
		return enums.LoginPortRU
	case enums.TR:
		return enums.LoginPortTR
	default:
		return ""
	}
}
