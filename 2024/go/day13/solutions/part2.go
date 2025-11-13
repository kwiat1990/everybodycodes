package solutions

import (
	"strings"
)

func Part2(input string) int {
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
