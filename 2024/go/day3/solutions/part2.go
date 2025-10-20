package solutions

func Part2(input string) int {
	grid := newGrid(input)
	return grid.mine(false)
}
