package solutions

import (
	"strconv"
	"strings"
)

func Part2(input string) int {
	type step struct {
		direction string
		value     int
	}

	lines := strings.Split(input, "\n")
	commands := make([][]step, len(lines))

	for i, line := range lines {
		notes := strings.Split(line, ",")
		commands[i] = make([]step, len(notes))

		for j, note := range notes {
			direction := string(note[0])
			val, _ := strconv.Atoi(note[1:])
			commands[i][j] = step{direction, val}
		}
	}

	type point struct {
		x, y, z int
	}

	segments := map[point]struct{}{}

	for _, commandSegment := range commands {
		currentPos := point{0, 0, 0}
		for _, command := range commandSegment {
			for range command.value {
				switch command.direction {
				case "U":
					currentPos.z += 1
				case "D":
					currentPos.z -= 1
				case "R":
					currentPos.x += 1
				case "L":
					currentPos.x -= 1
				case "F":
					currentPos.y += 1
				case "B":
					currentPos.y -= 1
				}
				segments[currentPos] = struct{}{}
			}
		}
	}

	return len(segments)
}
