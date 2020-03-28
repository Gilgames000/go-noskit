package entities

import "github.com/gilgames000/go-noskit/enums"

// Player represents an actor which is controller by a person.
type Player struct {
	Actor
	Name   string
	Class  enums.Class
	Gender enums.Gender
}
