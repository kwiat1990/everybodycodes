package solutions

import (
	"strings"
)

func Part1(input string) int {
	var words []string
	sentence := ""

	for line := range strings.Lines(input) {
		if strings.HasPrefix(line, "WORDS:") {
			after, _ := strings.CutPrefix(line, "WORDS:")
			words = strings.Split(strings.TrimSpace(after), ",")
		}

		if line != "" {
			sentence = line
		}
	}

	sum := 0
	for _, word := range words {
		sum += strings.Count(sentence, word)
	}

	return sum
}
