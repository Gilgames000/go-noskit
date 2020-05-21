package errors

import (
	"fmt"

	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
)

type BazaarInteractionError struct {
	Msg string
}

func (e BazaarInteractionError) Error() string {
	return "error while interacting with the bazaar NPC: " + e.Msg
}

type NPCNotFound struct {
	NPCID int
}

func (e NPCNotFound) Error() string {
	return fmt.Sprintf("there is no NPC with ID %d on the current map", e.NPCID)
}

type ShopNotFound struct {
	ShopType enums.ShopType
}

func (e ShopNotFound) Error() string {
	return fmt.Sprintf("there is no Shop with shop type %d on the current map", e.ShopType)
}

type ItemNotFound struct {
	VNum int
}

func (e ItemNotFound) Error() string {
	return fmt.Sprintf("there is no item with VNum %d", e.VNum)
}

type CharacterOutOfRange struct {
	Action string
}

func (e CharacterOutOfRange) Error() string {
	return fmt.Sprintf("the character is out of the minimum range required to perform the requested action (%s)", e.Action)
}

type PointNotWalkable struct {
	Point entities.Point
}

func (e PointNotWalkable) Error() string {
	return fmt.Sprintf("can't walk to the point (%d, %d) for it is not walkable", e.Point.X, e.Point.Y)
}

type CharacterCannotMove struct {
}

func (e CharacterCannotMove) Error() string {
	return "the character is not allowed to move at the moment"
}

type WrongCharacterStatus struct {
	Action string
	Status string
}

func (e WrongCharacterStatus) Error() string {
	return fmt.Sprintf("in order to perform the requested action (%s) the character should be in the following status: %s", e.Action, e.Status)
}

type ConnectionTimedOut struct {
}

func (e ConnectionTimedOut) Error() string {
	return "connection timed out"
}

type LoginFailed struct {
	ErrorCode string
}

func (e LoginFailed) Error() string {
	return fmt.Sprintf("login failed with error code %s", e.ErrorCode)
}

type JoinGameFailed struct {
	Slot    int
	Message string
}

func (e JoinGameFailed) Error() string {
	return fmt.Sprintf("failed to join the game with character in slot %d: %s", e.Slot, e.Message)
}
