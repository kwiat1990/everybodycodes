package solutions

import (
	"math"
	"strconv"
	"strings"
)

func Part3(input string) int {
	type point struct {
		col, row int
	}

	targets := map[point]struct{}{}
	segments := [...]point{
		{1, 2},
		{1, 1},
		{1, 0},
	}

	for _, line := range strings.Split(input, "\n") {
		position := strings.Fields(line)
		col, _ := strconv.Atoi(position[0])
		row, _ := strconv.Atoi(position[1])
		pos := point{col, row}
		targets[pos] = struct{}{}
	}

	canHitTarget := func(start, target point, power int) bool {
		col := start.col + (2 * power)
		row := start.row + power
		delta := target.col - col

		return target.row == row-delta && target.col == col+delta
	}

	sum := 0
	power := 1

	for target := range targets {
		minSum := math.MaxInt
		for i, segment := range segments {
			segNum := len(segments) - i
			for j := target.col - target.row; j > 0; i-- {
				pos := point{target.col - j, target.row - j}
				if canHitTarget(segment, pos, power) {
					minSum = min(minSum, segNum*power)
					break
				}
			}
		}
	}

	//
	// for len(targets) > 0 {
	// 	for i, segment := range segments {
	// 		segNum := len(segments) - i
	// 		for target := range targets {
	// 			if canHitTarget(segment, target, power) {
	// 				sum += segNum * power
	// 				delete(targets, target)
	// 			}
	// 		}
	// 	}
	// 	power += 1
	// }

	return sum
}
