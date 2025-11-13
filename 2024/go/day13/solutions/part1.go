package solutions

import (
	"cmp"
	"math"
	"slices"
	"strings"
)

type point struct {
	col, row int
}

func (p point) neighbors() []point {
	dirs := [4]point{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	neighbors := make([]point, len(dirs))

	for i, dir := range dirs {
		neighbors[i] = point{p.col + dir.col, p.row + dir.row}
	}

	return neighbors
}

func Part1(input string) int {
	var start, end point
	grid := map[point]int{}

	for row, line := range strings.Split(input, "\n") {
		for col, x := range line {
			switch {
			case x == 'E':
				end = point{col, row}
				grid[end] = 0
			case x == 'S':
				start = point{col, row}
				grid[start] = 0
			case x >= '0' && x <= '9':
				grid[point{col, row}] = int(x - '0')
			}
		}
	}

	paths := dijkstra(start, grid)
	return paths[end]
}

func dijkstra(start point, grid map[point]int) map[point]int {
	distances := map[point]int{}
	visited := map[point]bool{}

	for pos := range grid {
		distances[pos] = math.MaxInt32
	}

	distances[start] = 0

	type item struct {
		point
		dist int
	}

	pq := []item{{start, 0}}

	for len(pq) > 0 {
		slices.SortFunc(pq, func(a, b item) int {
			return cmp.Compare(a.dist, b.dist)
		})

		curr := pq[0]
		pq = pq[1:]

		if visited[curr.point] {
			continue
		}

		visited[curr.point] = true

		for _, pos := range curr.point.neighbors() {
			if val, ok := grid[pos]; ok {
				newDist := distances[curr.point] + platformDiff(grid[curr.point], val) + 1
				if newDist < distances[pos] {
					distances[pos] = newDist
					pq = append(pq, item{pos, newDist})
				}
			}
		}
	}

	return distances
}

func platformDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	// from - to => diff, 10 - diff => min of both
	// 5 - 7 = 2 => 2, 10-2 = 8 => 2 - default
	// 9 - 0 = 2 => 9, 10-9 = 1 => 1 - wrapping portal
	wrap := 10 - diff
	return min(diff, wrap)
}
