package main

import (
	"testing"

	"github.com/kwiat1990/everybodycodes/day1/solutions"
)

func TestPart1(t *testing.T) {
	input := "ABBAC"
	want := 5

	if ans := solutions.Part1(input); ans != want {
		t.Fatalf("got = %d, want = %d", ans, want)
	}
}

func TestPart2(t *testing.T) {
	input := "AxBCDDCAxD"
	want := 28

	if ans := solutions.Part2(input); ans != want {
		t.Fatalf("got = %d, want = %d", ans, want)
	}
}

func TestPart3(t *testing.T) {
	input := "xBxAAABCDxCC"
	want := 30

	if ans := solutions.Part3(input); ans != want {
		t.Fatalf("got = %d, want = %d", ans, want)
	}
}
