package solutions

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `1,2,3

^_^ -.- ^,-
>.- ^_^ >.<
-_- -.- >.<
    -.^ ^_^
    >.>`
	want := ">.- -.- ^,-"

	if got := Part1(input); got != want {
		t.Fatalf("got = %s, want = %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `1,2,3

^_^ -.- ^,-
>.- ^_^ >.<
-_- -.- >.<
    -.^ ^_^
    >.>`
	want := "280014668134"

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
