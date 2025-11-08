package solutions

import (
	"strconv"
	"strings"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	brightnesses := make([]int, len(lines))

	for i, line := range lines {
		value, _ := strconv.Atoi(line)
		brightnesses[i] = value
	}

	dots := [...]int{10, 5, 3, 1}
	sum := 0

	for _, brightness := range brightnesses {
		rest := brightness
		for _, dot := range dots {
			beetles := 0
			for rest > 0 {
				if rest-dot < 0 {
					break
				}
				rest -= dot
				beetles += 1
			}
			sum += beetles
		}
	}

	return sum
}
