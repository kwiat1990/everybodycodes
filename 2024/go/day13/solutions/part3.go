package solutions

import (
	"math"
	"strings"
)

func Part3(input string) int {
	var end point
	var starts []point
	grid := map[point]int{}

	for row, line := range strings.Split(input, "\n") {
		for col, x := range line {
			pos := point{col, row}
			switch {
			case x == 'E':
				end = point{col, row}
				grid[end] = 0
			case x == 'S':
				starts = append(starts, pos)
				grid[pos] = 0
			case x >= '0' && x <= '9':
				grid[pos] = int(x - '0')
			}
		}
	}

	paths := dijkstra(end, grid)
	minPath := math.MaxInt
	for _, start := range starts {
		minPath = min(minPath, paths[start])
	}

	return minPath
}
