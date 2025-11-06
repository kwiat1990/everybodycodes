package solutions

import (
	"strconv"
)

func Part2(input string) int {
	nullpointer, _ := strconv.Atoi(input)
	acolytes := 1111
	availableBlocks := 20240000

	// For test
	// acolytes = 5
	// availableBlocks = 50
	// nullpointer = 3

	width := 1
	usedBlocks := 1
	thickness := 1

	for usedBlocks < availableBlocks {
		width += 2
		thickness = (thickness * nullpointer) % acolytes
		usedBlocks += thickness * width
	}

	return (usedBlocks - availableBlocks) * width
}
