package solutions

import (
	"slices"
	"strings"
)

func Part2(input string) int {
	type point struct {
		col, row int
	}

	targets := map[point]bool{}
	segments := [3]point{}

	lines := strings.Split(input, "\n")
	slices.Reverse(lines)

	for row, line := range lines[1:] {
		for col, value := range line {
			pos := point{col, row}
			switch {
			case value == 'T':
				targets[pos] = false
			case value == 'H':
				targets[pos] = true
			case value == 'A':
				segments[2] = pos
			case value == 'B':
				segments[1] = pos
			case value == 'C':
				segments[0] = pos
			}
		}
	}

	sum := 0
	power := 1

	canHitTarget := func(start, target point, power int) bool {
		col := start.col + (2 * power)
		row := start.row + power
		colDiff := target.col - col

		return target.row == row-colDiff && target.col == col+colDiff
	}

	for len(targets) > 0 {
		for i, segment := range segments {
			segNum := len(segments) - i
			for target, isHardRock := range targets {
				if canHitTarget(segment, target, power) {
					val := segNum * power
					if isHardRock {
						val *= 2
					}
					sum += val
					delete(targets, target)
				}
			}
		}
		power += 1
	}

	return sum
}
