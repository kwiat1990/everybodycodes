package solutions

import (
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	nums := make([]int, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}

	slices.Sort(nums)
	smallest := nums[0]
	strikes := 0

	for _, num := range nums[1:] {
		strikes += num - smallest
	}

	return strikes
}
