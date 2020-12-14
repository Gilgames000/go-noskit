package pathfinder

import (
	"fmt"
	"github.com/gilgames000/go-noskit/entities"
	"testing"
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
	{entities.Point{0, 0}, entities.Point{4, 3}, mapAllWalkable, true},
	{entities.Point{2, 3}, entities.Point{3, 5}, mapAllWalkable, false},
	{entities.Point{4, 2}, entities.Point{4, 2}, mapAllWalkable, true},
	{entities.Point{0, 0}, entities.Point{4, 3}, mapOuterWalls, false},
	{entities.Point{2, 3}, entities.Point{3, 5}, mapOuterWalls, false},
	{entities.Point{4, 2}, entities.Point{4, 2}, mapOuterWalls, true},
	{entities.Point{0, 0}, entities.Point{4, 3}, mapHorizontalSplit, false},
	{entities.Point{3, 1}, entities.Point{0, 1}, mapHorizontalSplit, true},
	{entities.Point{4, 1}, entities.Point{4, 1}, mapHorizontalSplit, true},
	{entities.Point{0, 0}, entities.Point{4, 3}, mapWithObstacle, true},
	{entities.Point{3, 1}, entities.Point{0, 1}, mapWithObstacle, true},
	{entities.Point{4, 1}, entities.Point{4, 1}, mapWithObstacle, true},
	{entities.Point{0, 0}, entities.Point{4, 3}, mapNotWalkable, false},
	{entities.Point{2, 3}, entities.Point{3, 5}, mapNotWalkable, false},
	{entities.Point{4, 2}, entities.Point{4, 2}, mapNotWalkable, true},
}

func TestCalculatePath(t *testing.T) {
	pf := New()
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
