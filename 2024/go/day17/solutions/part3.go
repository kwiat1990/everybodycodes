package solutions

import (
	"math"
	"slices"
	"strings"
)

func Part3(input string) int {
	var nodes []point

	for row, line := range strings.Split(input, "\n") {
		for col, val := range line {
			if val == '*' {
				nodes = append(nodes, point{col, row})
			}
		}
	}

	prim := func() []int {
		var allMst []int
		mst := []point{nodes[0]}
		visited := map[point]bool{nodes[0]: true}
		minDistances := map[point]int{}
		distance := 0

		for _, node := range nodes[1:] {
			minDistances[node] = nodes[0].distance(node)
		}

		for len(visited) < len(nodes) {
			minNextDist := math.MaxInt
			var nextNode point

			for node, dist := range minDistances {
				if dist < minNextDist {
					minNextDist = dist
					nextNode = node
				}
			}

			visited[nextNode] = true
			delete(minDistances, nextNode)

			if minNextDist < 6 {
				mst = append(mst, nextNode)
				distance += minNextDist

				for candidate := range minDistances {
					newDist := nextNode.distance(candidate)
					minDistances[candidate] = min(minDistances[candidate], newDist)
				}
			} else {
				allMst = append(allMst, len(mst)+distance)
				distance = 0
				mst = mst[:0]
				mst = append(mst, nextNode)

				for candidate := range minDistances {
					minDistances[candidate] = nextNode.distance(candidate)
				}
			}
		}

		if len(mst) > 0 {
			allMst = append(allMst, len(mst)+distance)
		}

		return allMst
	}

	distances := prim()
	slices.Sort(distances)

	sum := 1
	for _, distance := range distances[len(distances)-3:] {
		sum *= distance
	}

	return sum
}
