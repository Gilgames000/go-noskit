package entities

// Thing represents an entity which can be found on the map (e.g. players, items, etc.).
// Why Thing and not Entity? Because we all love DooM :)
type Thing struct {
	ID       int
	Position Point
}
