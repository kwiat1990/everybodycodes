package solutions

import (
	"slices"
	"strings"
)

func Part2(input string) string {
	const trackTemplate2 = `S-=++=-==++=++=-=+=-=+=+=--=-=++=-==++=-+=-=+=-=+=+=++=-+==++=++=-=-=--
-                                                                     -
=                                                                     =
+                                                                     +
=                                                                     +
+                                                                     =
=                                                                     =
-                                                                     -
--==++++==+=+++-=+=-=+=-+-=+-=+-=+=-=+=--=+++=++=+++==++==--=+=++==+++-`

	var racePlans []racePlan

	for _, line := range strings.Split(input, "\n") {
		items := strings.Split(string(line), ":")
		rawScores := strings.Split(items[1], ",")
		actions := make([]int, len(rawScores))
		for i, score := range rawScores {
			actions[i] = getScore(score)
		}
		plan := racePlan{name: items[0], actions: actions, power: 0}
		racePlans = append(racePlans, plan)
	}

	rawTrack := strings.Split(trackTemplate2, "\n")
	var track []int

	// top row
	for i := 1; i < len(rawTrack[0]); i++ {
		item := string(rawTrack[0][i])
		if item != " " {
			track = append(track, getScore(item))
		}
	}

	// right side
	for i := 1; i < len(rawTrack)-1; i++ {
		lastIdx := len(rawTrack[i]) - 1
		item := string(rawTrack[i][lastIdx])
		if item != " " {
			track = append(track, getScore(item))
		}
	}

	// bottom row
	bottomRow := rawTrack[len(rawTrack)-1]
	for i := len(bottomRow) - 1; i >= 0; i-- {
		item := string(bottomRow[i])
		if item != " " {
			track = append(track, getScore(item))
		}
	}

	// left side
	for i := len(rawTrack) - 2; i > 0; i-- {
		item := string(rawTrack[i][0])
		if item != " " {
			track = append(track, getScore(item))
		}
	}

	track = append(track, 0)

	for i := range racePlans {
		prevPower := 10
		for j := range len(track) * 10 {
			trackIdx := j % len(track)
			action := track[trackIdx]

			idx := j % len(racePlans[i].actions)
			if action == 0 {
				action = racePlans[i].actions[idx]
			}

			currentPower := prevPower + action
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
