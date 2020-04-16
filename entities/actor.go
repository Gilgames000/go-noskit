package entities

// Actor represents a thing that is able to perform actions (e.g. players, NPCs, etc.).
type Actor struct {
	Thing
	CombatLevel  int
	HPPercentage int
	MPPercentage int
}
