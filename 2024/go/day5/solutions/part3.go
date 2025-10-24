package solutions

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func Part3(input string) int {
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

	var builder strings.Builder
	cache := map[string]struct{}{}
	highest := 0

	for i := 0; i < math.MaxInt; i++ {
		state := serialize(grid)
		if _, exists := cache[state]; exists {
			return highest
		}

		cache[state] = struct{}{}

		clapperIdx := i % 4
		clapper := grid[clapperIdx][0]

		targetCol := (clapperIdx + 1) % 4
		targetColLen := len(grid[targetCol])

		insertIdx := int(math.Abs(float64(clapper%(targetColLen*2)) - 1))
		if insertIdx > targetColLen {
			insertIdx = (targetColLen * 2) - insertIdx
		}

		grid[targetCol] = slices.Insert(grid[targetCol], insertIdx, clapper)
		grid[clapperIdx] = slices.Delete(grid[clapperIdx], 0, 1)

		builder.Reset()
		for j := 0; j < len(grid); j++ {
			firstEl := grid[j][0]
			builder.WriteString(strconv.Itoa(firstEl))
		}

		num, _ := strconv.Atoi(builder.String())
		if highest < num {
			highest = num
		}
	}

	return 0
}

func serialize(grid [][]int) string {
	var builder strings.Builder
	for _, col := range grid {
		for _, n := range col {
			builder.WriteByte(byte(n))
		}
	}
	return builder.String()
}
