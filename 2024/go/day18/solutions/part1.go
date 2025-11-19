package solutions

import (
	"slices"
	"strings"
)

func Part1(input string) int {
	type point struct {
		col, row int
	}

	grid := map[point]string{}
	var start point

	for row, line := range strings.Split(input, "\n") {
		for col, value := range line {
			if value == '#' {
				continue
			}

			pos := point{col, row}
			grid[pos] = string(value)

			if col == 0 && value == '.' {
				start = pos
			}
		}
	}

	bfs := func() int {
		type queue struct {
			point
			steps int
		}
		stack := []queue{{start, 0}}
		steps := 0

		var visited []point
		dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

		for len(stack) > 0 {
			current := stack[0]
			stack = stack[1:]

			if grid[current.point] == "P" {
				steps = max(steps, current.steps)
			}

			for _, dir := range dirs {
				next := point{current.col + dir.col, current.row + dir.row}

				if _, exists := grid[next]; exists && !slices.Contains(visited, next) {
					stack = append(stack, queue{next, current.steps + 1})
					visited = append(visited, next)
				}
			}
		}

		return steps
	}

	return bfs()
}
