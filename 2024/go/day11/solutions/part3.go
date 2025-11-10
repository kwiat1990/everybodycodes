package solutions

import (
	"math"
	"strings"
)

func Part3(input string) int {
	lines := strings.Split(input, "\n")
	conversions := map[string][]string{}

	for _, line := range lines {
		group := strings.Split(line, ":")
		individual := group[0]
		generation := strings.Split(group[1], ",")
		conversions[individual] = generation
	}

	minPop, maxPop := math.MaxInt, 0

	for individual := range conversions {
		population := map[string]int{individual: 1}

		for range 20 {
			nextDayPopulation := make(map[string]int)
			for current, count := range population {
				for _, next := range conversions[current] {
					nextDayPopulation[next] += count
				}
			}
			population = nextDayPopulation
		}

		sum := 0
		for _, count := range population {
			sum += count
		}

		minPop = min(minPop, sum)
		maxPop = max(maxPop, sum)
	}

	return maxPop - minPop
}
