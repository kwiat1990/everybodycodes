package solutions

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `#####.#####
#.........#
#.######.##
#.........#
###.#.#####
#H.......H#
###########`
	want := 26

	if got := Part1(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `##########.##########
#...................#
#.###.##.###.##.#.#.#
#..A#.#..~~~....#A#.#
#.#...#.~~~~~...#.#.#
#.#.#.#.~~~~~.#.#.#.#
#...#.#.B~~~B.#.#...#
#...#....BBB..#....##
#C............#....C#
#####################`
	want := 38

	if got := Part2(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart3(t *testing.T) {
	input := ``
	want := 0

	if got := Part3(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}
