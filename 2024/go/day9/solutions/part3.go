package solutions

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func Part3(input string) int {
	lines := strings.Split(input, "\n")
	amounts := make([]int, len(lines))

	for i, line := range lines {
		value, _ := strconv.Atoi(line)
		amounts[i] = value
	}

	coins := [...]int{1, 3, 5, 10, 15, 16, 20, 24, 25, 30, 37, 38, 49, 50, 74, 75, 100, 101}
	sum := 0

	for _, amount := range amounts {
		dp := slices.Repeat([]int{math.MaxInt}, amount+1)
		dp[0] = 0

		for i := 1; i <= amount; i++ {
			for _, coin := range coins {
				diff := i - coin
				if diff < 0 {
					break
				}
				dp[i] = min(dp[i], 1+dp[diff])
			}
		}
		sum += dp[amount]
	}

	return sum
}
