package solutions

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `**PCBS**
**RLNW**
BV....PT
CR....HZ
FL....JW
SG....MN
**FTZV**
**GMJH**`
	want := "PTBVRCZHFLJWGMNS"

	if got := Part1(input); got != want {
		t.Fatalf("got = %s, want = %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `**PCBS**
**RLNW**
BV....PT
CR....HZ
FL....JW
SG....MN
**FTZV**
**GMJH**`
	want := "1851"

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
