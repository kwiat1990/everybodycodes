package solutions

func Part3(input string) int {
	grid := newGrid(input)
	return grid.mine(true)
}
