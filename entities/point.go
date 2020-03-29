package entities

import "math"

// Point represents a point on the map.
type Point struct {
	X int
	Y int
}

// DistanceFrom returns the distance between two points.
func (p1 *Point) DistanceFrom(p2 Point) int {
	lx := p1.X - p2.X
	ly := p1.Y - p2.Y
	sqDist := lx*lx + ly*ly
	dist := math.Sqrt(float64(sqDist))

	return int(dist)
}
