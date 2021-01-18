package pathfinder_test

import (
	"fmt"
	"testing"

	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/pathfinder"
)

var mapAllWalkable = [][]bool{
	{true, true, true, true},
	{true, true, true, true},
	{true, true, true, true},
	{true, true, true, true},
	{true, true, true, true},
	{true, true, true, true},
}

var mapOuterWalls = [][]bool{
	{false, false, false, false},
	{false, true, true, false},
	{false, true, true, false},
	{false, true, true, false},
	{false, true, true, false},
	{false, false, false, false},
}

var mapHorizontalSplit = [][]bool{
	{true, true, false, true},
	{true, true, false, true},
	{true, true, false, true},
	{true, true, false, true},
	{true, true, false, true},
	{true, true, false, true},
}

var mapWithObstacle = [][]bool{
	{true, true, true, true},
	{true, true, false, true},
	{true, true, false, true},
	{true, true, false, true},
	{true, true, false, true},
	{true, true, true, true},
}

var mapNotWalkable = [][]bool{
	{false, false, false, false},
	{false, false, false, false},
	{false, false, false, false},
	{false, false, false, false},
	{false, false, false, false},
	{false, false, false, false},
}

var pathfinderTests = []struct {
	p1, p2          entities.Point
	walkabilityGrid [][]bool
	shouldWork      bool
}{
	{entities.Point{}, entities.Point{X: 4, Y: 3}, mapAllWalkable, true},
	{entities.Point{}, entities.Point{X: 4, Y: 3}, mapOuterWalls, false},
	{entities.Point{}, entities.Point{X: 4, Y: 3}, mapHorizontalSplit, false},
	{entities.Point{}, entities.Point{X: 4, Y: 3}, mapWithObstacle, true},
	{entities.Point{}, entities.Point{X: 4, Y: 3}, mapNotWalkable, false},
	{entities.Point{X: 3, Y: 1}, entities.Point{Y: 1}, mapWithObstacle, true},
	{entities.Point{X: 3, Y: 1}, entities.Point{Y: 1}, mapHorizontalSplit, true},
	{entities.Point{X: 2, Y: 3}, entities.Point{X: 3, Y: 5}, mapAllWalkable, false},
	{entities.Point{X: 4, Y: 2}, entities.Point{X: 4, Y: 2}, mapAllWalkable, true},
	{entities.Point{X: 2, Y: 3}, entities.Point{X: 3, Y: 5}, mapOuterWalls, false},
	{entities.Point{X: 4, Y: 2}, entities.Point{X: 4, Y: 2}, mapOuterWalls, true},
	{entities.Point{X: 4, Y: 1}, entities.Point{X: 4, Y: 1}, mapHorizontalSplit, true},
	{entities.Point{X: 4, Y: 1}, entities.Point{X: 4, Y: 1}, mapWithObstacle, true},
	{entities.Point{X: 2, Y: 3}, entities.Point{X: 3, Y: 5}, mapNotWalkable, false},
	{entities.Point{X: 4, Y: 2}, entities.Point{X: 4, Y: 2}, mapNotWalkable, true},
}

func TestCalculatePath(t *testing.T) {
	pf := pathfinder.New()
	for i, tt := range pathfinderTests {
		t.Run(fmt.Sprintf("calculate path test %d", i), func(t *testing.T) {
			_, err := pf.FindPath(tt.p1, tt.p2, tt.walkabilityGrid)
			if tt.shouldWork && err != nil {
				t.Errorf("pathfinding should've worked, but it failed\nmap: %v\np1: %v\n p2: %v\n",
					tt.walkabilityGrid, tt.p1, tt.p2)
			} else if !tt.shouldWork && err == nil {
				t.Errorf("pathfinding should've failed, but it worked\nmap: %v\np1: %v\n p2: %v\n",
					tt.walkabilityGrid, tt.p1, tt.p2)
			}
		})
	}
}
