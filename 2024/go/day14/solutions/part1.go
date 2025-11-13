package solutions

import (
	"strconv"
	"strings"
)

func Part1(input string) int {
	type step struct {
		direction string
		value     int
	}

	notes := strings.Split(input, ",")
	commands := make([]step, len(notes))

	for i, note := range notes {
		direction := string(note[0])
		val, _ := strconv.Atoi(note[1:])
		commands[i] = step{direction, val}
	}

	sum := 0
	highest := 0

	for _, command := range commands {
		switch command.direction {
		case "U":
			sum += command.value
		case "D":
			sum -= command.value
		}
		highest = max(highest, sum)
	}

	return highest
}
