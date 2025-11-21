package solutions

import (
	"slices"
	"strings"
)

func Part2(input string) int {
	type point struct{ x, y int }

	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")

	grid := map[point]int{}
	targets := []rune{'S', 'A', 'B', 'C'}
	var S, A, B, C point

	for y, line := range lines {
		for x, ch := range line {
			pos := point{x, y}

			switch ch {
			case '#':
				continue
			case '-':
				grid[pos] = -2
			case '+':
				grid[pos] = 1
			default:
				grid[pos] = -1
				if slices.Contains(targets, ch) {
					idx := slices.Index(targets, ch)
					switch idx {
					case 0:
						S = pos
					case 1:
						A = pos
					case 2:
						B = pos
					case 3:
						C = pos
					}
				}
			}
		}
	}

	advanceStage := func(stage int, p point) int {
		switch {
		case stage == 0 && p == A:
			return 1
		case stage == 1 && p == B:
			return 2
		case stage == 2 && p == C:
			return 3
		default:
			return stage
		}
	}

	type stateKey struct {
		pos   point
		dir   point
		stage int
	}

	type state struct {
		key      stateKey
		altitude int
		time     int
	}

	key := stateKey{S, point{0, 1}, 0}
	start := state{key, 10_000, 0}
	queue := []state{start}

	visited := map[stateKey]int{}
	visited[key] = 10_000

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.key.pos == S && cur.key.stage == 3 && cur.altitude >= 10_000 {
			return cur.time
		}

		dir := cur.key.dir
		candidates := []point{
			dir,
			{-dir.y, dir.x},
			{dir.y, -dir.x},
		}

		for _, nd := range candidates {
			next := point{cur.key.pos.x + nd.x, cur.key.pos.y + nd.y}
			if _, ok := grid[next]; !ok {
				continue
			}
			altitude := cur.altitude + grid[next]
			nextStage := advanceStage(cur.key.stage, next)
			nextKey := stateKey{pos: next, dir: nd, stage: nextStage}

			if visited[nextKey] < altitude {
				visited[nextKey] = altitude
				queue = append(queue, state{
					key:      nextKey,
					altitude: altitude,
					time:     cur.time + 1,
				})
			}
		}
	}

	return -1
}
