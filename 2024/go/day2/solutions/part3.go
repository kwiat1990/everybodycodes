package solutions

import (
	"strings"
)

type Grid struct {
	cells  [][]byte
	height int
	width  int
}

func (g Grid) wrapIndex(i int) int {
	return (i%g.width + g.width) % g.width
}

func newGrid(cells [][]byte) Grid {
	height := len(cells)
	width := len(cells[0])
	return Grid{
		cells,
		height,
		width,
	}
}

type point struct {
	row, col int
}

type searchBuffer struct {
	backward []byte
	forward  []byte
	points   []point
}

func (sb *searchBuffer) clear() {
	sb.backward = sb.backward[:0]
	sb.forward = sb.forward[:0]
	sb.points = sb.points[:0]
}

func newSearchBuffer(length int) *searchBuffer {
	return &searchBuffer{
		backward: make([]byte, length),
		forward:  make([]byte, length),
		points:   make([]point, length),
	}

}

func collectCells(word []byte, candidates []byte, points []point, lookup [][]bool) {
	for i := range word {
		if word[i] != candidates[i] {
			return
		}
	}

	for _, c := range points {
		lookup[c.row][c.col] = true
	}
}

func Part3(input string) int {
	lines := strings.Split(input, "\n\n")
	first, _ := strings.CutPrefix(lines[0], "WORDS:")

	var words [][]byte
	maxWordLen := 0
	for word := range strings.SplitSeq(first, ",") {
		word = strings.TrimSpace(word)
		wordBytes := []byte(word)
		words = append(words, wordBytes)
		if len(wordBytes) > maxWordLen {
			maxWordLen = len(wordBytes)
		}

	}

	var sentences [][]byte
	for sentence := range strings.SplitSeq(lines[1], "\n") {
		sentences = append(sentences, []byte(sentence))
	}

	grid := newGrid(sentences)

	lookup := make([][]bool, grid.height)
	for i := range len(lookup) {
		lookup[i] = make([]bool, grid.width)
	}

	buffer := newSearchBuffer(maxWordLen)

	for _, word := range words {
		for row := 0; row < grid.height; row++ {
			for col := 0; col < grid.width; col++ {
				buffer.clear()

				for i := range len(word) {
					currentCol := grid.wrapIndex(col + i)
					buffer.forward = append(buffer.forward, grid.cells[row][currentCol])
					buffer.points = append(buffer.points, point{row, currentCol})
				}

				collectCells(word, buffer.forward, buffer.points, lookup)

				for i := len(buffer.forward) - 1; i >= 0; i-- {
					buffer.backward = append(buffer.backward, buffer.forward[i])
				}

				collectCells(word, buffer.backward, buffer.points, lookup)
			}
		}

		for col := 0; col < grid.width; col++ {
			for row := 0; row <= grid.height-len(word); row++ {
				buffer.clear()

				for i := range len(word) {
					currentRow := row + i
					buffer.forward = append(buffer.forward, grid.cells[currentRow][col])
					buffer.points = append(buffer.points, point{currentRow, col})
				}

				collectCells(word, buffer.forward, buffer.points, lookup)

				for i := len(buffer.forward) - 1; i >= 0; i-- {
					buffer.backward = append(buffer.backward, buffer.forward[i])
				}

				collectCells(word, buffer.backward, buffer.points, lookup)
			}
		}
	}

	sum := 0
	for _, row := range lookup {
		for _, col := range row {
			if col {
				sum += 1
			}
		}
	}

	return sum
}
