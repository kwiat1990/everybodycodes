package solutions

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `2
4
7
16`
	want := 10

	if got := Part1(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `33
41
55
99`
	want := 10

	if got := Part2(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart3(t *testing.T) {
	input := `156488
352486
546212`
	want := 10449

	if got := Part3(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}
