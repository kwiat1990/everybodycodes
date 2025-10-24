package solutions

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `2 3 4 5
3 4 5 2
4 5 2 3
5 2 3 4`
	want := 2323

	if got := Part1(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `2 3 4 5
6 7 8 9`
	want := 50877075

	if got := Part2(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart3(t *testing.T) {
	input := `2 3 4 5
6 7 8 9`
	want := 6584

	if got := Part3(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}
