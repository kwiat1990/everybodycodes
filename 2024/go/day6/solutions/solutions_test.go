package solutions

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `RR:A,B,C
A:D,E
B:F,@
C:G,H
D:@
E:@
F:@
G:@
H:@`
	want := "RRB@"

	if got := Part1(input); got != want {
		t.Fatalf("got = %s, want = %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `RR:A,B,C
A:D,E
B:F,@
C:G,H
D:@
E:@
F:@
G:@
H:@`
	want := "RB@"

	if got := Part2(input); got != want {
		t.Fatalf("got = %s, want = %s", got, want)
	}
}

func TestPart3(t *testing.T) {
	input := `RR:A,B,C
A:D,E
B:F,@
C:G,H
D:@
E:@
F:@
G:@
H:@`
	want := "RB@"

	if got := Part3(input); got != want {
		t.Fatalf("got = %s, want = %s", got, want)
	}
}
