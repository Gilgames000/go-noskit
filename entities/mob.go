package entities

// Mob represents an anctor controlled by the computer. It can be attacked
// and is generally regarded as an enemy.
type Mob struct {
	Actor
	VNum int
}
