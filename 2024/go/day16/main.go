package main

import (
	"fmt"
	"os"

	"github.com/kwiat1990/everybodycodes/day16/solutions"
)

func main() {
	funcs := []func(string) string{solutions.Part1, solutions.Part2, solutions.Part3}
	files := []string{"input1.txt", "input2.txt", "input3.txt"}

	for i, fn := range funcs {
		file, err := os.ReadFile(fmt.Sprintf("inputs/%s", files[i]))

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("part%d: %s\n", i+1, fn(string(file)))
	}
}
