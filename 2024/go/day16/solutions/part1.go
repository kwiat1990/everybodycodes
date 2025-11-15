package solutions

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) string {

	parts := strings.Split(input, "\n\n")
	positions := strings.Split(parts[0], ",")
	multiplies := make([]int, len(positions))

	for i, x := range positions {
		num, _ := strconv.Atoi(x)
		multiplies[i] = num
	}

	lines := strings.Split(parts[1], "\n")
	width := len(strings.Fields(lines[0]))

	symbols := make([][]string, width)

	re := regexp.MustCompile(`(.{3})(?: |$)`)

	for _, line := range strings.Split(parts[1], "\n") {
		matches := re.FindAllString(line, -1)

		for i, match := range matches {
			match = strings.TrimSpace(match)
			if match == "" {
				continue
			}
			idx := i % width
			symbols[idx] = append(symbols[idx], match)
		}
	}

	result := make([]string, len(symbols))

	for i, line := range symbols {
		idx := 100 * multiplies[i] % len(line)
		result[i] = line[idx]
	}

	return strings.Join(result, " ")
}
