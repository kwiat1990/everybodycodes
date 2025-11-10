package solutions

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `A:B,C
B:C,A
C:A`
	want := 8

	if got := Part1(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `A:B,C
B:C,A
C:A
Z:A`
	want := 89

	if got := Part2(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart3(t *testing.T) {
	input := `A:B,C
B:C,A,A
C:A`
	want := 268815

	if got := Part3(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}
