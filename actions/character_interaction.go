package actions

import (
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
	"github.com/gilgames000/go-noskit/errors"
)

// BazaarInteractor lets you interact with the logged-in character.
type CharacterInteractor struct {
	character        CharacterGateway
	currentMap       MapGateway
	characterManager CharacterManagementGateway
}

func NewCharacterInteractor(character CharacterGateway, currentMap MapGateway, characterManager CharacterManagementGateway) *CharacterInteractor {
	return &CharacterInteractor{
		character:        character,
		currentMap:       currentMap,
		characterManager: characterManager,
	}
}

// CharacterGateway provides an abstraction over low-level methods used to
// perform actions with the logged-in character.
type CharacterGateway interface {
	Info() entities.Character
	WalkTo(p entities.Point) error
	CanMove() bool
	CanAttack() bool
}

// CharacterManagementGateway provides an abstraction over low-level methods used to
// manage the characters on the logged-in account.
type CharacterManagementGateway interface {
	JoinGame(slot int) error
	CurrentStatus() enums.CharacterStatus
}

// MapGateway provides information and methods related to the structure
// of the current map.
type MapGateway interface {
	Info() entities.Map
	DistanceBetween(p1 entities.Point, p2 entities.Point) (int, error)
	FindPath(p1 entities.Point, p2 entities.Point) ([]entities.Point, error)
	IsWalkable(p entities.Point) bool
}

// NPCGateway provides information about the NPCs present on the
// current map.
type NPCGateway interface {
	All() []entities.NPC
	SearchByID(id int) (entities.NPC, bool)
}

// ShopGateway provides information about the shops present on the
// current map.
type ShopGateway interface {
	All() []entities.Shop
	SearchByID(id int) (entities.Shop, bool)
	SearchByShopType(st enums.ShopType) (entities.Shop, bool)
}

// WalkTo will make you character walk to the specified position
// on the map. It will return an error if the character is unable
// to move or the specified position is not walkable/reachable.
func (ci *CharacterInteractor) WalkTo(p entities.Point) error {
	if !ci.currentMap.IsWalkable(p) {
		return &errors.PointNotWalkable{Point: p}
	}

	if !ci.character.CanMove() {
		return &errors.CharacterCannotMove{}
	}

	return ci.character.WalkTo(p)
}

// JoinGame joins the game with the character in the selected slot.
func (ci *CharacterInteractor) JoinGame(slot int) error {
	if ci.characterManager.CurrentStatus() != enums.CharacterSelection {
		return errors.WrongCharacterStatus{
			Action: "join game",
			Status: "character selection",
		}
	}

	return ci.characterManager.JoinGame(slot)
}
