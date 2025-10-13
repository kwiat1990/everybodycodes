package solutions

func Part3(input string) int {
	sum := 0
	mapping := map[rune]int{
		'A': 0,
		'B': 1,
		'C': 3,
		'D': 5,
	}

	boost := map[int]int{
		0: 0,
		1: 0,
		2: 2,
		3: 6,
	}

	for i := 2; i < len(input); i += 3 {
		arr := []byte{input[i-2], input[i-1], input[i]}
		boostCount := 0

		for _, r := range arr {
			if r != 'x' {
				boostCount += 1
			}
			sum += mapping[rune(r)]
		}

		sum += boost[boostCount]
	}

	return sum
}
