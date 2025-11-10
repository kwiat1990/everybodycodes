package solutions

import (
	"slices"
	"strings"
)

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	conversions := map[string][]string{}

	for _, line := range lines {
		group := strings.Split(line, ":")
		individual := group[0]
		generation := strings.Split(group[1], ",")
		conversions[individual] = generation
	}

	population := []string{"Z"}

	for range 10 {
		tmp := slices.Clone(population)
		population = population[:0]
		for _, individual := range tmp {
			population = append(population, conversions[individual]...)
		}
	}

	return len(population)
}
