package solutions

import (
	"slices"
	"strconv"
	"strings"
)

func Part3(input string) int {
	lines := strings.Split(input, "\n")
	nums := make([]int, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}

	slices.Sort(nums)
	middle := nums[len(nums)/2]
	strikes := 0

	for _, num := range nums {
		diff := num - middle
		if diff < 0 {
			diff = -diff
		}
		strikes += diff
	}

	return strikes
}
