package solutions

import (
	"slices"
	"strings"
)

func Part1(input string) string {
	grid := make([][]string, 8)

	for row, line := range strings.Split(input, "\n") {
		grid[row] = make([]string, 8)
		for col, value := range strings.Split(line, "") {
			grid[row][col] = value
		}
	}

	var builder strings.Builder
	for i := 2; i < 6; i++ {
		for j := 2; j < 6; j++ {
			row := grid[i]
			col := make([]string, 8)
			for x, r := range grid {
				col[x] = r[j]
			}

			for _, ch := range row {
				if ch != "." && slices.Contains(col, ch) {
					grid[i][j] = ch
					builder.WriteString(ch)
					break
				}
			}
		}
	}
	return builder.String()
}
