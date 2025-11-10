package solutions

import (
	"strconv"
	"strings"
)

func Part2(input string) string {
	chunks := strings.Split(input, "\n\n")
	sum := 0

	for _, chunk := range chunks {
		lines := strings.Split(chunk, "\n")
		samplesCount := len(strings.Fields(lines[0]))

		grids := make([][8][8]string, samplesCount)

		for j, samples := range lines {
			for i, sample := range strings.Fields(samples) {
				for k, ch := range sample {
					grids[i][j][k] = string(ch)
				}
			}
		}

		for _, grid := range grids {
			var builder strings.Builder

			for row := 2; row < 6; row++ {
				for col := 2; col < 6; col++ {
					var curCol [8]string
					for x := 0; x < 8; x++ {
						curCol[x] = grid[x][col]
					}

				outer:
					for _, rowVal := range grid[row] {
						for _, colVal := range curCol {
							if rowVal == "." {
								continue
							}
							if rowVal == colVal {
								grid[row][col] = rowVal
								builder.WriteString(rowVal)
								break outer
							}
						}
					}
				}
			}

			for i, ch := range builder.String() {
				val := int(ch - 'A' + 1)
				sum += (i + 1) * val
			}
		}
	}

	return strconv.Itoa(sum)
}
