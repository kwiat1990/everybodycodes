package solutions

import (
	"strings"
)

func Part1(input string) int {
	type point struct {
		col, row int
	}

	grid := map[point]int{}
	var start point

	for row, line := range strings.Split(input, "\n") {
		for col, val := range line {
			pos := point{col, row}
			switch val {
			case 'S':
				start = pos
			case '.':
				grid[pos] = -1
			case '-':
				grid[pos] = -2
			case '+':
				grid[pos] = 1
			}
		}
	}

	type state struct {
		point
		direction point
		time      int
		altitude  int
	}

	bfs := func() int {
		queue := []state{
			{start, point{0, 1}, 0, 1000},
		}

		type stateKey struct {
			pos  point
			dir  point
			time int
		}

		visited := map[stateKey]int{}
		maxAltitude := 0

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			if current.time == 100 {
				maxAltitude = max(maxAltitude, current.altitude)
				continue
			}

			dirs := []point{
				current.direction,
				{current.direction.row, -current.direction.col},
				{-current.direction.row, current.direction.col},
			}

			for _, dir := range dirs {
				pos := point{current.col + dir.col, current.row + dir.row}

				if val, exists := grid[pos]; exists {
					next := state{pos, dir, current.time + 1, current.altitude + val}
					key := stateKey{next.point, next.direction, next.time}

					if visited[key] < next.altitude {
						queue = append(queue, next)
						visited[key] = next.altitude
					}
				}
			}
		}

		return maxAltitude
	}

	return bfs()
}
