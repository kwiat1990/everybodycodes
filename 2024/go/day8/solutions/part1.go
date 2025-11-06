package solutions

import (
	"slices"
	"strconv"
)

func Part1(input string) int {
	availableBlocks, _ := strconv.Atoi(input)
	var blocks []int
	blocksUsed := 0

	for level := 0; blocksUsed <= availableBlocks; level++ {
		newBlocks := level + 1
		if level > 0 {
			newBlocks = blocks[level-1] + 2
		}
		blocks = append(blocks, newBlocks)
		blocksUsed += newBlocks
	}

	return (blocksUsed - availableBlocks) * slices.Max(blocks)

}
