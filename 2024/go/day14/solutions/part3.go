package solutions

import (
	"math"
	"strconv"
	"strings"
)

func Part3(input string) int {
	type step struct {
		direction string
		value     int
	}

	type point struct {
		x, y, z int
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

	leaves := make([]point, len(lines))
	segments := map[point]struct{}{}

	for i, commandSegment := range commands {
		currentPos := point{0, 0, 0}
		for j, command := range commandSegment {
			for range command.value {
				switch command.direction {
				case "U":
					currentPos.y += 1
				case "D":
					currentPos.y -= 1
				case "R":
					currentPos.x += 1
				case "L":
					currentPos.x -= 1
				case "F":
					currentPos.z += 1
				case "B":
					currentPos.z -= 1
				}
				segments[currentPos] = struct{}{}
				if j == len(commandSegment)-1 {
					leaves[i] = currentPos
				}
			}
		}
	}

	bfs := func(start point) map[point]int {
		distances := map[point]int{start: 0}
		queue := []point{start}

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			neighbors := []point{
				{current.x + 1, current.y, current.z},
				{current.x - 1, current.y, current.z},
				{current.x, current.y + 1, current.z},
				{current.x, current.y - 1, current.z},
				{current.x, current.y, current.z + 1},
				{current.x, current.y, current.z - 1},
			}

			for _, neighbor := range neighbors {
				if _, exists := segments[neighbor]; exists {
					if _, visited := distances[neighbor]; !visited {
						distances[neighbor] = distances[current] + 1
						queue = append(queue, neighbor)
					}
				}
			}
		}

		return distances
	}

	minMurkiness := math.MaxInt
	for pos := range segments {
		if pos.x != 0 {
			// we're not on the main trunk x=0
			continue
		}

		distances := bfs(pos)
		murkiness := 0
		for _, leaf := range leaves {
			murkiness += distances[leaf]
		}
		minMurkiness = min(minMurkiness, murkiness)
	}

	return minMurkiness
}
