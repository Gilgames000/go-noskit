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

type NPCNotFoundError struct {
	NPCID int
}

func (e NPCNotFoundError) Error() string {
	return fmt.Sprintf("there is no NPC with ID %d on the current map", e.NPCID)
}

type ShopNotFoundError struct {
	ShopType enums.ShopType
}

func (e ShopNotFoundError) Error() string {
	return fmt.Sprintf("there is no Shop with shop type %d on the current map", e.ShopType)
}

type ItemNotFoundError struct {
	VNum int
}

func (e ItemNotFoundError) Error() string {
	return fmt.Sprintf("there is no item with VNum %d", e.VNum)
}

type OutOfRangeError struct {
	Msg string
}

func (e OutOfRangeError) Error() string {
	return "character out of range to perform action: " + e.Msg
}

type PointNotWalkableError struct {
	Point entities.Point
}

func (e PointNotWalkableError) Error() string {
	return fmt.Sprintf("can't walk to the point (%d, %d) for it is not walkable", e.Point.X, e.Point.Y)
}

type CharacterCannotMoveError struct {
}

func (e CharacterCannotMoveError) Error() string {
	return "the character is not allowed to move at the moment"
}
