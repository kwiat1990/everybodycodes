package solutions

import (
	"math"
	"strings"
)

type point struct {
	col, row int
}

func (p point) distance(other point) int {
	return abs(p.col-other.col) + abs(p.row-other.row)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Part1(input string) int {
	var nodes []point

	for row, line := range strings.Split(input, "\n") {
		for col, val := range line {
			if val == '*' {
				nodes = append(nodes, point{col, row})
			}
		}
	}

	prim := func() int {
		distance := 0
		candidateDistances := map[point]int{}
		mst := []point{nodes[0]}
		visited := map[point]bool{nodes[0]: true}

		for len(visited) < len(nodes) {

			for _, current := range mst {
				if !visited[current] {
					continue
				}

				for _, candidate := range nodes {
					if visited[candidate] {
						continue
					}

					if dist, exists := candidateDistances[candidate]; !exists {
						candidateDistances[candidate] = current.distance(candidate)
					} else {
						candidateDistances[candidate] = min(dist, current.distance(candidate))
					}
				}
			}

			minNextDist := math.MaxInt
			var nextNode point

			for key, value := range candidateDistances {
				if value < minNextDist {
					minNextDist = value
					nextNode = key
				}
			}

			mst = append(mst, nextNode)
			clear(candidateDistances)
			distance += minNextDist
			visited[nextNode] = true
		}

		return distance
	}

	return len(nodes) + prim()
}
