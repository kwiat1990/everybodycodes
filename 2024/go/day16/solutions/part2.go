package solutions

import (
	"regexp"
	"strconv"
	"strings"
)

func Part2(input string) string {
	parts := strings.Split(input, "\n\n")
	positions := strings.Split(parts[0], ",")
	multipliers := make([]int, len(positions))

	for i, raw := range positions {
		num, _ := strconv.Atoi(raw)
		multipliers[i] = num
	}

	re := regexp.MustCompile(`(.{3})(?: |$)`)
	lines := strings.Split(parts[1], "\n")
	width := len(strings.Fields(lines[0]))

	symbolColumns := make([][]string, width)

	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		for i, match := range matches {
			match = strings.TrimSpace(match)
			if match == "" {
				continue
			}
			col := i % width
			symbolColumns[col] = append(symbolColumns[col], match)
		}
	}

	// Compute the full-cycle period via LCM
	lcmPeriod := 1
	for _, column := range symbolColumns {
		lcmPeriod = lcm(lcmPeriod, len(column))
	}

	// Break the huge number into:
	// - some full cycles
	// - and one partial cycle
	totalPulls := 202420242024
	fullCycles := totalPulls / lcmPeriod
	partialCyclePulls := totalPulls % lcmPeriod

	// Simulate the partial cycle first
	coinsInPartialCycle := spin(0, partialCyclePulls, symbolColumns, multipliers)

	// Simulate the rest of a full LCM cycle (starting after partial)
	coinsInOneFullCycle := spin(partialCyclePulls, lcmPeriod, symbolColumns, multipliers)

	// Total coins per cycle
	coinsPerFullCycle := coinsInPartialCycle + coinsInOneFullCycle

	// Multiply by the number of full cycles
	coinsFromFullCycles := fullCycles * coinsPerFullCycle

	return strconv.Itoa(coinsInPartialCycle + coinsFromFullCycles)
}

func spin(startPull, endPull int, symbolColumns [][]string, multipliers []int) int {
	coins := 0
	sequence := make([]string, len(symbolColumns))

	for pull := startPull + 1; pull <= endPull; pull++ {
		for colIndex, column := range symbolColumns {
			idx := pull * multipliers[colIndex] % len(column)
			sequence[colIndex] = column[idx]
		}
		coins += addWonCoins(sequence)
	}
	return coins
}

func addWonCoins(sequence []string) int {
	frequencies := map[rune]int{}
	for _, str := range sequence {
		frequencies[rune(str[0])] += 1
		frequencies[rune(str[2])] += 1
	}

	total := 0
	for _, val := range frequencies {
		if val > 2 {
			total += val - 2
		}
	}
	return total
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
