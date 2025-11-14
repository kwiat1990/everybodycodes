package solutions

import (
	"fmt"
	"strings"
)

// Part2 it's the Travelling Politician Problem, a variation of the Travelling Salesman Problem (TSP).
func Part2(input string) int {
	type point struct {
		col, row int
	}

	type herb struct {
		point
		value string
	}
	grid := map[point]string{}
	var herbs []herb
	var entry herb

	for row, line := range strings.Split(input, "\n") {
		for col, value := range line {
			if value == '#' || value == '~' {
				continue
			}

			pos := point{col, row}
			if row == 0 && value == '.' {
				entry = herb{pos, string(value)}
				continue
			}

			if 'A' <= value && value <= 'Z' {
				herbs = append(herbs, herb{pos, string(value)})
			}

			grid[pos] = string(value)
		}
	}

	fmt.Println(herbs)

	dfs := func(start point) map[point]int {
		distances := map[point]int{}

		var explore func(current point, dist int)
		explore = func(current point, dist int) {
			// If we've seen this node before with a shorter distance, skip
			if prevDist, exists := distances[current]; exists && prevDist <= dist {
				return
			}

			distances[current] = dist

			neighbors := []point{
				{current.col + 1, current.row},
				{current.col - 1, current.row},
				{current.col, current.row + 1},
				{current.col, current.row - 1},
			}

			for _, neighbor := range neighbors {
				if _, exists := grid[neighbor]; exists {
					explore(neighbor, dist+1)
				}
			}
		}

		explore(start, 0)
		return distances
	}

	// bfs := func(start, end point) int {
	// 	distances := map[point]int{start: 0}
	// 	queue := []point{start}
	//
	// 	for len(queue) > 0 {
	// 		current := queue[0]
	// 		queue = queue[1:]
	//
	// 		if current == end {
	// 			fmt.Println(distances)
	// 			return distances[current]
	// 		}
	//
	// 		neighbors := []point{
	// 			{current.col + 1, current.row},
	// 			{current.col - 1, current.row},
	// 			{current.col, current.row + 1},
	// 			{current.col, current.row - 1},
	// 		}
	//
	// 		for _, neighbor := range neighbors {
	// 			if _, exists := grid[neighbor]; exists {
	// 				if _, visited := distances[neighbor]; !visited {
	// 					distances[neighbor] = distances[current] + 1
	// 					queue = append(queue, neighbor)
	// 				}
	// 			}
	// 		}
	// 	}
	//
	// 	return math.MaxInt
	// }

	sum := 0
	start := entry
	for _, herb := range herbs {

		dist := dfs(start.point)
		sum += dist[herb.point]
		start = herb
	}
	fmt.Println(sum)
	dist := dfs(start.point)
	sum += dist[entry.point]
	fmt.Println(start.point, entry.point)
	return sum
}
