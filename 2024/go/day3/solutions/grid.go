package solutions

import (
	"math"
	"slices"
	"strings"
)

type point struct {
	col, row int
}

func (p point) add(other point) point {
	return point{p.col + other.col, p.row + other.row}
}

type grid struct {
	positions map[point]int
}

func (g grid) canDig(pos point, includeDiagonal bool) bool {
	neighbors := g.getNeighbors(pos, includeDiagonal)
	requiredNeighborsCount := 4

	if includeDiagonal {
		requiredNeighborsCount = 8
	}

	if len(neighbors) < requiredNeighborsCount {
		return false
	}

	for _, n := range neighbors {
		if math.Abs(float64(g.positions[n]-g.positions[pos])) > 1 {
			return false
		}
	}

	return true
}

func (g grid) getNeighbors(pos point, includeDiagonal bool) []point {
	dirs := []point{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	if includeDiagonal {
		diagonal := []point{
			{1, 1},
			{-1, 1},
			{1, -1},
			{-1, -1},
		}
		dirs = append(dirs, diagonal...)
	}

	neighbors := make([]point, len(dirs))
	for i, dir := range dirs {
		neighbors[i] = pos.add(dir)
	}
	return neighbors
}

func (g grid) mine(includeDiagonal bool) int {
	var stack []point
	sum := 0

	for len(g.positions) > 1 {
		for pos, _ := range g.positions {
			if g.canDig(pos, includeDiagonal) {
				stack = append(stack, pos)
			}
		}

		for pos := range g.positions {
			if slices.Contains(stack, pos) {
				g.positions[pos] += 1
			} else {
				delete(g.positions, pos)
			}
		}

		sum += len(stack)
		stack = stack[:0]
	}

	return sum
}

func newGrid(input string) grid {
	positions := map[point]int{}

	for row, line := range strings.Split(input, "\n") {
		for col, val := range strings.Split(line, "") {
			pos := point{col, row}
			if val == "#" {
				positions[pos] = 1
			}
		}
	}

	return grid{positions}
}
