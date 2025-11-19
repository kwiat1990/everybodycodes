package solutions

import (
	"math"
	"strings"
)

func Part3(input string) int {
	type point struct {
		col, row int
	}

	grid := map[point]string{}

	for row, line := range strings.Split(input, "\n") {
		for col, value := range line {
			if value == '#' {
				continue
			}

			pos := point{col, row}
			grid[pos] = string(value)
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

		visited := map[point]bool{}
		distances := map[point]int{
			start: 0,
		}

		dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

		for len(stack) > 0 {
			current := stack[0]
			stack = stack[1:]

			for _, dir := range dirs {
				next := point{current.col + dir.col, current.row + dir.row}

				if _, exists := grid[next]; exists && !visited[next] {
					stack = append(stack, queue{next, current.steps + 1})
					distances[next] = current.steps + 1
					visited[next] = true
				}
			}
		}

		return distances
	}

	distances := map[point]int{}
	for key, val := range grid {
		if val == "P" {
			for k, v := range bfs(key) {
				distances[k] += v
			}
		}
	}

	minSum := math.MaxInt
	for _, v := range distances {
		minSum = min(minSum, v)
	}

	return minSum
}
