package solutions

import (
	"slices"
	"strings"
)

type racePlan struct {
	name    string
	actions []int
	power   int
}

func Part1(input string) string {
	var racePlans []racePlan

	for _, line := range strings.Split(input, "\n") {
		items := strings.Split(line, ":")
		rawScores := strings.Split(items[1], ",")
		actions := make([]int, len(rawScores))
		for i, score := range rawScores {
			actions[i] = getScore(score)
		}
		plan := racePlan{name: items[0], actions: actions, power: 0}
		racePlans = append(racePlans, plan)
	}

	for i := range racePlans {
		prevPower := 10
		for j := range 10 {
			idx := j % len(racePlans[i].actions)
			currentPower := prevPower + racePlans[i].actions[idx]
			racePlans[i].power += currentPower
			prevPower = currentPower
		}
	}

	slices.SortFunc(racePlans, func(a, b racePlan) int {
		switch {
		case a.power < b.power:
			return 1
		case a.power > b.power:
			return -1
		default:
			return 0
		}
	})

	var builder strings.Builder
	for _, plan := range racePlans {
		builder.WriteString(plan.name)
	}

	return builder.String()
}

func getScore(sign string) int {
	switch sign {
	case "+":
		return 1
	case "-":
		return -1
	default:
		return 0
	}
}
