package solutions

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `U5,R3,D2,L5,U4,R5,D2`
	want := 7

	if got := Part1(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `U5,R3,D2,L5,U4,R5,D2
U6,L1,D2,R3,U2,L1`
	want := 32

	if got := Part2(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart3a(t *testing.T) {
	input := `U20,L1,B1,L2,B1,R2,L1,F1,U1
U10,F1,B1,R1,L1,B1,L1,F1,R2,U1
U30,L2,F1,R1,B1,R1,F2,U1,F1
U25,R1,L2,B1,U1,R2,F1,L2
U16,L1,B1,L1,B3,L1,B1,F1`
	want := 46

	if got := Part3(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart3b(t *testing.T) {
	input := `U5,R3,D2,L5,U4,R5,D2
U6,L1,D2,R3,U2,L1`
	want := 5

	if got := Part3(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}
