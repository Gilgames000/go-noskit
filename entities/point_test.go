package entities

import (
	"strconv"
	"testing"
)

var pointDistanceFromTests = []struct {
	p1, p2   Point
	distance int
}{
	{
		p1:       Point{0, 0},
		p2:       Point{0, 0},
		distance: 0,
	},
	{
		p1:       Point{3, 4},
		p2:       Point{3, 4},
		distance: 0,
	},
	{
		p1:       Point{2, 7},
		p2:       Point{6, 6},
		distance: 4,
	},
	{
		p1:       Point{-1, -1},
		p2:       Point{-1, -1},
		distance: 0,
	},
	{
		p1:       Point{-4, 5},
		p2:       Point{9, -2},
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
