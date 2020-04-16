package entities

// Character represents the character controlled by the user.
type Character struct {
	Player
	Speed     int
	CurrentHP int
	CurrentMP int
	MaxHP     int
	MaxMP     int
}
