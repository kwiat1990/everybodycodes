package solutions

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := "ABBAC"
	want := 5

	if ans := Part1(input); ans != want {
		t.Fatalf("got = %d, want = %d", ans, want)
	}
}

func TestPart2(t *testing.T) {
	input := "AxBCDDCAxD"
	want := 28

	if ans := Part2(input); ans != want {
		t.Fatalf("got = %d, want = %d", ans, want)
	}
}

func TestPart3(t *testing.T) {
	input := "xBxAAABCDxCC"
	want := 30

	if ans := Part3(input); ans != want {
		t.Fatalf("got = %d, want = %d", ans, want)
	}
}
