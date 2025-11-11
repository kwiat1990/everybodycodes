package solutions

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `.............
.C...........
.B......T....
.A......T.T..
=============`
	want := 13

	if got := Part1(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `.............
.C...........
.B......H....
.A......T.H..
=============`
	want := 22

	if got := Part2(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart3(t *testing.T) {
	input := `6 5
6 7
10 5`
	want := 11

	if got := Part3(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}
