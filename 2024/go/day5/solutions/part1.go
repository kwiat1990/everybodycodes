package solutions

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	grid := make([][]int, 0, len(lines))

	for _, line := range lines {
		fields := strings.Fields(line)

		for len(grid) < len(fields) {
			grid = append(grid, make([]int, 0, len(lines)))
		}

		for col, ch := range fields {
			num, _ := strconv.Atoi(ch)
			grid[col] = append(grid[col], num)
		}
	}

	for i := 0; i < 10; i++ {
		clapperIdx := i % 4
		clapper := grid[clapperIdx][0]

		targetCol := (clapperIdx + 1) % 4
		targetColLen := len(grid[targetCol])

		insertIdx := int(math.Abs(float64((clapper % (targetColLen * 2)) - 1)))
		if insertIdx > targetColLen {
			insertIdx = (targetColLen * 2) - insertIdx
		}

		grid[targetCol] = slices.Insert(grid[targetCol], insertIdx, clapper)
		grid[clapperIdx] = slices.Delete(grid[clapperIdx], 0, 1)
	}

	result := 0
	for i := 0; i < len(grid); i++ {
		firstEl := grid[i][0]
		result = result*10 + firstEl
	}
	return result
}
