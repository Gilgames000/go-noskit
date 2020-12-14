package pathfinder

import (
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/errors"
	"github.com/gilgames000/go-noskit/gamestate"

	"github.com/SolarLune/paths"
)

var _ gamestate.Pathfinder = &Pathfinder{}

type Pathfinder struct {
}

func New() *Pathfinder {
	return &Pathfinder{}
}

func (p Pathfinder) FindPath(p1, p2 entities.Point, walkabilityGrid [][]bool) ([]entities.Point, error) {
	if p1 == p2 {
		return []entities.Point{}, nil
	}

	path := calculatePath(p1, p2, walkabilityGrid)
	if len(path.Cells) == 0 {
		return nil, errors.NoPathToPoint{
			From: p1,
			To:   p2,
		}
	}

	return cellsToPoints(path.Cells), nil
}

func (p Pathfinder) DistanceBetween(p1, p2 entities.Point, walkabilityGrid [][]bool) (int, error) {
	if p1 == p2 {
		return 0, nil
	}

	path := calculatePath(p1, p2, walkabilityGrid)
	if len(path.Cells) == 0 {
		return -1, errors.NoPathToPoint{
			From: p1,
			To:   p2,
		}
	}

	return len(path.Cells), nil
}

func calculatePath(p1, p2 entities.Point, walkabilityGrid [][]bool) *paths.Path {
	grid := fillGrid(walkabilityGrid)
	if p1.X > grid.Width()-1 || p1.Y > grid.Height()-1 || p2.X > grid.Width() || p2.Y > grid.Height() {
		return &paths.Path{}
	}
	return grid.GetPath(
		grid.Get(p1.X, p1.Y),
		grid.Get(p2.X, p2.Y),
		true,
	)
}

func fillGrid(walkabilityGrid [][]bool) *paths.Grid {
	h := len(walkabilityGrid[0])
	w := len(walkabilityGrid)
	grid := paths.NewGrid(w, h)

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if walkabilityGrid[i][j] {
				grid.Get(i, j).Cost = 1
				grid.Get(i, j).Walkable = true
			} else {
				grid.Get(i, j).Walkable = false
			}
		}
	}

	return grid
}

func cellsToPoints(cells []*paths.Cell) []entities.Point {
	path := make([]entities.Point, len(cells))
	for i := range cells {
		path[i] = entities.Point{
			X: cells[i].X,
			Y: cells[i].Y,
		}
	}

	return path
}
