package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gilgames000/go-noskit/actions"
	"github.com/gilgames000/go-noskit/data"
	"github.com/gilgames000/go-noskit/datastore"
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
	"github.com/gilgames000/go-noskit/gamestate"
	"github.com/gilgames000/go-noskit/gfclient"
	"github.com/gilgames000/go-noskit/packets/parser"
	packetsrv "github.com/gilgames000/go-noskit/packets/server"
	"github.com/gilgames000/go-noskit/pathfinder"
	"github.com/gilgames000/go-noskit/sockets"
)

type DummyGameClientGateway struct {
}

func (d *DummyGameClientGateway) Version() string {
	return os.Getenv("NOSTALE_CLIENT_VERSION")
}

func (d *DummyGameClientGateway) Hash() string {
	return os.Getenv("NOSTALE_CLIENT_HASH")
}

func registerPackets(packetParser *parser.NosPacketParser) {
	packetParser.RegisterPacket(packetsrv.NsTeST{})
	packetParser.RegisterPacket(packetsrv.ConnectionFailure{})
	packetParser.RegisterPacket(packetsrv.CharacterListItem{})
	packetParser.RegisterPacket(packetsrv.Fail{})
	packetParser.RegisterPacket(packetsrv.CharacterListEnd{})
	packetParser.RegisterPacket(packetsrv.OK{})
	packetParser.RegisterPacket(packetsrv.Info{})
	packetParser.RegisterPacket(packetsrv.CharacterInfo{})
	packetParser.RegisterPacket(packetsrv.CharacterStatus{})
	packetParser.RegisterPacket(packetsrv.CharacterLevel{})
	packetParser.RegisterPacket(packetsrv.CharacterPosition{})
	packetParser.RegisterPacket(packetsrv.EntityCondition{})
	packetParser.RegisterPacket(packetsrv.BazaarSearchResults{})
	packetParser.RegisterPacket(packetsrv.NPCRequest{})
	packetParser.RegisterPacket(packetsrv.SpawnMob{})
	packetParser.RegisterPacket(packetsrv.SpawnNPC{})
	packetParser.RegisterPacket(packetsrv.Shop{})
	packetParser.RegisterPacket(packetsrv.NPCInfo{})
	packetParser.RegisterPacket(packetsrv.WindowOpen{})
}

func getCountryID(lang string) enums.CountryID {
	switch lang {
	case "en":
		return enums.EN
	case "de":
		return enums.DE
	case "fr":
		return enums.FR
	case "it":
		return enums.IT
	case "pl":
		return enums.PL
	case "es":
		return enums.ES
	case "cz":
		return enums.CZ
	case "ru":
		return enums.RU
	case "tr":
		return enums.TR
	default:
		return enums.CountryID(-1)
	}
}

func main() {
	packetParser := parser.New()
	registerPackets(packetParser)

	pf := pathfinder.New()

	mapDataStore := datastore.NewMapDataStore(
		data.NewRawMapLoader(os.Getenv("NOSTALE_MAPS_DIRECTORY")),
	)

	itemDataStore, err := datastore.NewItemDataStore(
		data.NewCSVItemsLoader(
			os.Getenv("NOSTALE_ITEMS_CSV_PATH"),
			os.Getenv("NOSTALE_ITEMS_CSV_HAS_HEADER") == "true",
		),
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	gameSocket := sockets.NewGameSocket(packetParser)

	user := actions.User{
		Email:            os.Getenv("NOSTALE_EMAIL"),
		Username:         os.Getenv("NOSTALE_USERNAME"),
		Password:         os.Getenv("NOSTALE_PASSWORD"),
		InstallationUUID: os.Getenv("NOSTALE_INSTALL_UUID"),
		Locale:           os.Getenv("NOSTALE_LOCALE"),
	}

	userInteractor := actions.NewUserInteractor(
		gamestate.NewUserGateway(
			&gfclient.GFClient{},
			sockets.NewLoginSocket(packetParser),
			gameSocket,
			&DummyGameClientGateway{},
		),
	)

	characterInteractor := actions.NewCharacterInteractor(
		gamestate.NewCharacterGateway(gameSocket, mapDataStore, pf),
		gamestate.NewMapGateway(gameSocket, mapDataStore, pf),
		gamestate.NewCharacterManagementGateway(gameSocket),
	)

	bazaarInteractor := actions.NewBazaarInteractor(
		gamestate.NewItemGateway(itemDataStore),
		gamestate.NewBazaarGateway(gameSocket, itemDataStore),
		gamestate.NewCharacterGateway(gameSocket, mapDataStore, pf),
		gamestate.NewMapGateway(gameSocket, mapDataStore, pf),
		gamestate.NewNPCGateway(gameSocket),
		gamestate.NewShopGateway(gameSocket),
	)

	lang := os.Getenv("NOSTALE_SERVER_LANG")
	servers, err := userInteractor.Login(user, lang, getCountryID(lang))
	if err != nil {
		fmt.Printf("Login error: %s\n", err.Error())
		os.Exit(-1)
	}

	characters, err := userInteractor.Connect(servers[1].Channels[2])
	if err != nil {
		fmt.Printf("Connect error: %s\n", err.Error())
		os.Exit(-1)
	}

	time.Sleep(1250 * time.Millisecond)
	err = characterInteractor.JoinGame(characters[0].Slot)
	if err != nil {
		fmt.Printf("JoinGame error: %s\n", err.Error())
		os.Exit(-1)
	}

	time.Sleep(5 * time.Second)
	err = characterInteractor.WalkTo(entities.Point{
		X: 9,
		Y: 28,
	})
	if err != nil {
		fmt.Printf("Walk error: %s\n", err.Error())
		os.Exit(-1)
	}

	err = bazaarInteractor.Open()
	if err != nil {
		fmt.Printf("Bazaar error: %s\n", err.Error())
		os.Exit(-1)
	}

	res, err := bazaarInteractor.SearchItemByVNum(2282)
	if err != nil {
		fmt.Printf("Bazaar error: %s\n", err.Error())
		os.Exit(-1)
	}

	fmt.Println(res)

	select {}
}
