package solutions

import (
	"math"
	"strings"
)

func Part2(input string) int {
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
		minDistances := map[point]int{}
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

					if dist, exists := minDistances[candidate]; !exists {
						minDistances[candidate] = current.distance(candidate)
					} else {
						minDistances[candidate] = min(dist, current.distance(candidate))
					}
				}
			}

			minNextDist := math.MaxInt
			var nextNode point

			for key, value := range minDistances {
				if value < minNextDist {
					minNextDist = value
					nextNode = key
				}
			}

			mst = append(mst, nextNode)
			clear(minDistances)
			distance += minNextDist
			visited[nextNode] = true
		}

		return distance
	}

	return len(nodes) + prim()
}
