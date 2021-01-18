package entities_test

import (
	"strconv"
	"testing"

	"github.com/gilgames000/go-noskit/entities"
)

var pointDistanceFromTests = []struct {
	p1, p2   entities.Point
	distance int
}{
	{
		p1:       entities.Point{},
		p2:       entities.Point{},
		distance: 0,
	},
	{
		p1:       entities.Point{X: 3, Y: 4},
		p2:       entities.Point{X: 3, Y: 4},
		distance: 0,
	},
	{
		p1:       entities.Point{X: 2, Y: 7},
		p2:       entities.Point{X: 6, Y: 6},
		distance: 4,
	},
	{
		p1:       entities.Point{X: -1, Y: -1},
		p2:       entities.Point{X: -1, Y: -1},
		distance: 0,
	},
	{
		p1:       entities.Point{X: -4, Y: 5},
		p2:       entities.Point{X: 9, Y: -2},
		distance: 14,
	},
}

func TestPointDistanceFrom(t *testing.T) {
	for i, tt := range pointDistanceFromTests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if d := tt.p1.DistanceFrom(tt.p2); d != tt.distance {
				t.Errorf("expected: %d\nfound: %d", tt.distance, d)
			}
		})
	}
}
