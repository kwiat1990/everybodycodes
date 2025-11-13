package solutions

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `#######
#6769##
S50505E
#97434#
#######`
	want := 28

	if got := Part1(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `#######
#6769##
S50505E
#97434#
#######`
	want := 28

	if got := Part2(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}

func TestPart3(t *testing.T) {
	input := `SSSSSSSSSSS
S674345621S
S###6#4#18S
S53#6#4532S
S5450E0485S
S##7154532S
S2##314#18S
S971595#34S
SSSSSSSSSSS`
	want := 14

	if got := Part3(input); got != want {
		t.Fatalf("got = %d, want = %d", got, want)
	}
}
