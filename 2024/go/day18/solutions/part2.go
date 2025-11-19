package solutions

import (
	"slices"
	"strings"
)

func Part2(input string) int {
	type point struct {
		col, row int
	}

	targetsCount := 0
	grid := map[point]string{}
	var startA, startB point

	for row, line := range strings.Split(input, "\n") {
		for col, value := range line {
			if value == '#' {
				continue
			}

			pos := point{col, row}
			grid[pos] = string(value)

			if value == 'P' {
				targetsCount += 1
			}

			if col == 0 && value == '.' {
				startA = pos
			}

			if col == len(line)-1 && value == '.' {
				startB = pos
			}
		}
	}

	bfs := func(start point) map[point]int {
		type queue struct {
			point
			steps int
		}

		stack := []queue{
			{start, 0},
		}
		distances := map[point]int{}
		var visited []point

		dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

		for len(stack) > 0 {
			currentFront := stack[0]
			stack = stack[1:]

			if grid[currentFront.point] == "P" {
				distances[currentFront.point] = currentFront.steps
			}

			if len(distances) == targetsCount {
				break
			}

			for _, dir := range dirs {
				next := point{currentFront.col + dir.col, currentFront.row + dir.row}

				if _, exists := grid[next]; exists && !slices.Contains(visited, next) {
					stack = append(stack, queue{next, currentFront.steps + 1})
					visited = append(visited, next)
				}
			}
		}

		return distances
	}

	front := bfs(startA)
	back := bfs(startB)
	maxPath := 0

	for pos, val := range back {
		minPath := min(val, front[pos])
		maxPath = max(maxPath, minPath)
	}

	return maxPath
}
