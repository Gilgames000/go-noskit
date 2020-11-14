package actions

import "github.com/gilgames000/go-noskit/enums"

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
	user UserGateway
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
}

// AccountCharacter holds the information about a character on the logged-in
// user account.
type AccountCharacter struct {
	Slot int
	Name string
}

func (ui *UserInteractor) Login(user User, serverLang string, countryID enums.CountryID) []GameServer {
	panic("implement me")
}
func (ui *UserInteractor) Connect(user User, serverLang string) []GameServer {
	panic("implement me")
}
