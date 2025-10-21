package solutions

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `..........
..###.##..
...####...
..######..
..######..
...####...
..........`
	want := 35

	if got := Part1(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `..........
..###.##..
...####...
..######..
..######..
...####...
..........`
	want := 35

	if got := Part2(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart3(t *testing.T) {
	input := `..........
..###.##..
...####...
..######..
..######..
...####...
..........`
	want := 29

	if got := Part3(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}
