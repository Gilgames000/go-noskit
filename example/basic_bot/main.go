package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

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

type GameClientGateway struct {
}

func (d *GameClientGateway) Version() string {
	return os.Getenv("NOSTALE_CLIENT_VERSION")
}

func (d *GameClientGateway) Hash() string {
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
	st := os.Getenv("NOSTALE_MAPS_DIRECTORY")
	fmt.Println(st)
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

	characterGateway := gamestate.NewCharacterGateway(gameSocket, mapDataStore, pf)
	mapGateway := gamestate.NewMapGateway(gameSocket, mapDataStore, pf)

	userInteractor := actions.NewUserInteractor(
		gamestate.NewUserGateway(
			&gfclient.GFClient{},
			sockets.NewLoginSocket(packetParser),
			gameSocket,
			&GameClientGateway{},
		),
	)

	characterInteractor := actions.NewCharacterInteractor(
		characterGateway,
		mapGateway,
		gamestate.NewCharacterManagementGateway(gameSocket),
	)

	bazaarInteractor := actions.NewBazaarInteractor(
		gamestate.NewItemGateway(itemDataStore),
		gamestate.NewBazaarGateway(gameSocket, itemDataStore),
		characterGateway,
		mapGateway,
		gamestate.NewNPCGateway(gameSocket),
		gamestate.NewShopGateway(gameSocket),
	)

	lang := os.Getenv("NOSTALE_SERVER_LANG")
	servers, err := userInteractor.Login(user, lang, getCountryID(lang))
	if err != nil {
		fmt.Printf("Login error: %s\n", err.Error())
		os.Exit(-1)
	}
	var s actions.GameServer
	if lang == "de" {

		if servers[0].Name == "Ancelloan" {
			s = servers[0]
		}
		if servers[1].Name == "Ancelloan" {
			s = servers[1]
		}
		if servers[2].Name == "Ancelloan" {
			s = servers[2]
		}
	} else if lang == "it" {
		if servers[0].Name == "Flare" {
			s = servers[0]
		}
		if servers[1].Name == "Flare" {
			s = servers[1]
		}
	}

	characters, err := userInteractor.Connect(s.Channels[2])
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
	time.Sleep(1 * time.Second)

	err = bazaarInteractor.Open()
	if err != nil {
		fmt.Printf("Bazaar error: %s\n", err.Error())
		os.Exit(-1)
	}
	time.Sleep(1 * time.Second)
	var db *sql.DB
	var insert *sql.Rows
	d := "d0321167:dummeskind@tcp(85.13.145.183)/d0321167"

	db, err = sql.Open("mysql", d)
	//Here starts the bazar lookup loop
	var items []int = []int{2282, 1030, 2283, 2196, 2284, 2285, 1012, 1011, 1013, 1014, 1029, 5060, 315, 316, 317, 318, 319, 320, 321, 322}

	for _, v := range items {

		query := "CREATE TABLE IF NOT EXISTS `" + strconv.Itoa(v) + "` ( Date datetime, Price int )"
		insert, err = db.Query(query)
		// if there is an error inserting, handle it
		if err != nil {
			panic(err.Error())
		}
		// be careful deferring Queries if you are using transactions
		defer insert.Close()
	}
	for {

		for _, v := range items {
			res, err := bazaarInteractor.SearchItemByVNum(v)
			if err != nil {
				fmt.Printf("Bazaar error: %s\n", err.Error())
				os.Exit(-1)
			}

			Price := strconv.Itoa(res[0].Price)
			VNUM := strconv.Itoa(v)
			fmt.Println(VNUM)
			fmt.Println(Price)
			fmt.Println("")

			if err != nil {
				log.Printf("Error %s when opening DB\n", err)
				return
			}

			t := time.Now().Format("2006-01-02 15:04:05")
			insert, err = db.Query("INSERT INTO `" + VNUM + "`(`Date`, `Price`) VALUES ( '" + t + "'," + Price + ")")
			// if there is an error inserting, handle it
			if err != nil {
				panic(err.Error())
			}
			// be careful deferring Queries if you are using transactions
			insert.Close()
			time.Sleep(5 * time.Second)

		}
		time.Sleep(60 * 15 * time.Second)
	}

}
