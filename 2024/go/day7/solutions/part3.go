package solutions

import (
	"fmt"
	"strings"
)

func Part3(input string) string {
	const trackTemplate3 = `S+= +=-== +=++=     =+=+=--=    =-= ++=     +=-  =+=++=-+==+ =++=-=-=--
- + +   + =   =     =      =   == = - -     - =  =         =-=        -
= + + +-- =-= ==-==-= --++ +  == == = +     - =  =    ==++=    =++=-=++
+ + + =     +         =  + + == == ++ =     = =  ==   =   = =++=
= = + + +== +==     =++ == =+=  =  +  +==-=++ =   =++ --= + =
+ ==- = + =   = =+= =   =       ++--          +     =   = = =--= ==++==
=     ==- ==+-- = = = ++= +=--      ==+ ==--= +--+=-= ==- ==   =+=    =
-               = = = =   +  +  ==+ = = +   =        ++    =          -
-               = + + =   +  -  = + = = +   =        +     =          -
--==++++==+=+++-= =-= =-+-=  =+-= =-= =--   +=++=+++==     -=+=++==+++-`

	items := strings.Split(input, ":")
	rawScores := strings.Split(items[1], ",")
	actions := make([]int, len(rawScores))

	for i, score := range rawScores {
		actions[i] = getScore(score)
	}

	// rivalPlan := racePlan{name: items[0], actions: actions, power: 0}

	rawTrack := strings.Split(trackTemplate3, "\n")
	// height, width := len(rawTrack), len(rawTrack[0])
	var track strings.Builder

	for row, line := range rawTrack {
		for col := range strings.Split(line, "") {
			for _, neighbor := range neighbors(col, row, len(rawTrack[row]), 10) {
				c, r := neighbor[0], neighbor[1]
				// if c >= len(rawTrack[r]) {
				// 	fmt.Println("oko", r, row, len(line), len(rawTrack[row]), len(rawTrack[r]))
				// 	continue
				// }
				cell := rawTrack[r][c]
				if cell == ' ' {
					continue
				}
				if cell == 'S' {
					continue
				}
				track.WriteByte(cell)
			}
		}
	}
	track.WriteByte('S')
	fmt.Println(track.String())

	segment := make([]int, 11)
	for range 11 {
		// prevPower := 10
		for range 11 {

			// if action == 0 {
			// 	action = plan.actions[idx]
			// }
			//
			// currentPower := prevPower + action
			// racePlans[i].power += currentPower
			// prevPower = currentPower
		}
		segment = segment[:0]
	}

	// TODO implement solution
	return ""
}

func neighbors(col, row, width, height int) [][2]int {
	dirs := [][2]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}
	var n [][2]int
	for _, dir := range dirs {
		c, r := dir[0]+col, dir[1]+row
		if 0 <= c && c < width && 0 <= r && r < height {
			n = append(n, [2]int{c, r})
		}
	}
	return n
}
