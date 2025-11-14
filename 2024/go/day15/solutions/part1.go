package solutions

import (
	"math"
	"strings"
)

func Part1(input string) int {
	type point struct {
		col, row int
	}

	grid := map[point]bool{}
	var start point

	for row, line := range strings.Split(input, "\n") {
		for col, value := range line {
			pos := point{col, row}
			if value == '#' {
				continue
			}
			if row == 0 && value == '.' {
				start = pos
				continue
			}
			if value == '.' {
				grid[pos] = false
			}
			if value == 'H' {
				grid[pos] = true
			}
		}
	}

	bfs := func() int {
		distances := map[point]int{start: 0}
		queue := []point{start}

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			if grid[current] {
				return distances[current]
			}

			neighbors := []point{
				{current.col + 1, current.row},
				{current.col - 1, current.row},
				{current.col, current.row + 1},
				{current.col, current.row - 1},
			}

			for _, neighbor := range neighbors {
				if _, exists := grid[neighbor]; exists {
					if _, visited := distances[neighbor]; !visited {
						distances[neighbor] = distances[current] + 1
						queue = append(queue, neighbor)
					}
				}
			}
		}

		return math.MaxInt
	}

	return bfs() * 2
}
