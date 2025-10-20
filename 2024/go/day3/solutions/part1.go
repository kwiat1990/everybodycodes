package solutions

func Part1(input string) int {
	grid := newGrid(input)
	return grid.mine(false)
}
