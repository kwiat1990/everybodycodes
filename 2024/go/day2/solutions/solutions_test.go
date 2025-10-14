package solutions

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `WORDS:THE,OWE,MES,ROD,HER

	AWAKEN THE POWER ADORNED WITH THE FLAMES BRIGHT IRE`
	want := 4

	if ans := Part1(input); ans != want {
		t.Fatalf("got = %d, want = %d", ans, want)
	}
}

func TestPart2(t *testing.T) {
	input := `WORDS:THE,OWE,MES,ROD,HER,QAQ

	AWAKEN THE POWE ADORNED WITH THE FLAMES BRIGHT IRE
	THE FLAME SHIELDED THE HEART OF THE KINGS
	POWE PO WER P OWE R
	THERE IS THE END
	QAQAQ`
	want := 42

	if ans := Part2(input); ans != want {
		t.Fatalf("got = %d, want = %d", ans, want)
	}
}

func TestPart3(t *testing.T) {
	input := `WORDS:THE,OWE,MES,ROD,RODEO

	HELWORLT
	ENIGWDXL
	TRODEOAL`
	want := 10

	if ans := Part3(input); ans != want {
		t.Fatalf("got = %d, want = %d", ans, want)
	}
}
