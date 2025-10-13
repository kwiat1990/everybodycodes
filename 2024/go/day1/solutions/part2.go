package solutions

func Part2(input string) int {
	sum := 0
	mapping := map[rune]int{
		'A': 0,
		'B': 1,
		'C': 3,
		'D': 5,
	}

	for i := 1; i < len(input); i += 2 {
		a, b := input[i-1], input[i]
		sum += mapping[rune(a)] + mapping[rune(b)]

		if a != 'x' && b != 'x' {
			sum += 2
		}
	}

	return sum
}
