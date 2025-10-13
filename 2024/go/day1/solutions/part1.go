package solutions

func Part1(input string) int {
	sum := 0
	for _, r := range input {
		switch r {
		case 'B':
			sum += 1
		case 'C':
			sum += 3
		}
	}
	return sum
}
