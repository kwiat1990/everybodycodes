package solutions

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `A:+,-,=,=
B:+,=,-,+
C:=,-,+,+
D:=,=,=,+`
	want := "BDCA"

	if got := Part1(input); got != want {
		t.Fatalf("got = %s, want = %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `A:+,-,=,=
B:+,=,-,+
C:=,-,+,+
D:=,=,=,+

S+===
-   +
=+=-+`
	want := "DCBA"

	if got := Part2(input); got != want {
		t.Fatalf("got = %s, want = %s", got, want)
	}
}

func TestPart3(t *testing.T) {
	input := ``
	want := ""

	if got := Part3(input); got != want {
		t.Fatalf("got = %s, want = %s", got, want)
	}
}
