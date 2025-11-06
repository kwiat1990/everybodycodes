package solutions

import (
	"strconv"
)

func Part3(input string) int {
	nullpointer, _ := strconv.Atoi(input)
	acolytes := 10
	availableBlocks := 202400000

	// For test
	// nullpointer = 2
	// acolytes = 5
	// availableBlocks = 160

	usedBlocks := 1
	currentLayer := []int{1}
	width := 1
	thickness := 1

	for usedBlocks < availableBlocks {
		width += 2
		thickness = (thickness*nullpointer)%acolytes + acolytes
		usedBlocks += thickness * width

		tmp := make([]int, width)
		tmp[0] = thickness
		for i, val := range currentLayer {
			tmp[i+1] = val + thickness
		}
		tmp[len(tmp)-1] = thickness
		currentLayer = tmp
	}

	redundantBlocks := 0
	for _, val := range currentLayer[1 : len(currentLayer)-1] {
		redundantBlocks += (2 * len(currentLayer)) * val % acolytes
	}

	return (usedBlocks - redundantBlocks) - availableBlocks
}
