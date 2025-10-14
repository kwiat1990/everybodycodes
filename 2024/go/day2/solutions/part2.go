package solutions

import (
	"slices"
	"strings"
)

func Part2(input string) int {
	var words []string
	var sentences [][]string

	for line := range strings.Lines(input) {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "WORDS:") {
			after, _ := strings.CutPrefix(line, "WORDS:")
			words = strings.Split(after, ",")
		} else if line != "" {
			sentences = append(sentences, strings.Split(line, ""))
		}
	}

	sum := 0
	for _, sentence := range sentences {
		mapper := map[int]struct{}{}

		for _, word := range words {
			for start, end := 0, len(word); end <= len(sentence); start, end = start+1, end+1 {
				chunk := sentence[start:end]
				backwardChunk := make([]string, len(chunk))

				copy(backwardChunk, chunk)
				slices.Reverse(backwardChunk)

				if word == strings.Join(chunk, "") || word == strings.Join(backwardChunk, "") {
					for i := range len(word) {
						mapper[start+i] = struct{}{}
					}
				}
			}
		}

		sum += len(mapper)
	}

	return sum
}
