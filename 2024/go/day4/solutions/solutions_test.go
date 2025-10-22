package solutions

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `3
4
7
8`
	want := 10

	if got := Part1(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `3
4
7
8`
	want := 10

	if got := Part2(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart3(t *testing.T) {
	input := `2
4
5
6
8`
	want := 8

	if got := Part3(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}
