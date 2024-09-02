package matrix_graph

func closedIsland(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	var noOfClosedIsland int

	var dfs func(row, col int) bool
	dfs = func(row, col int) bool { // is closed island
		if grid[row][col] == 1 {
			return true
		}
		if (row == 0 || col == 0 || row == rows-1 || col == cols-1) && grid[row][col] == 0 {
			return false
		}
		grid[row][col] = 1

		top := dfs(row-1, col)
		down := dfs(row+1, col)
		left := dfs(row, col-1)
		right := dfs(row, col+1)

		return top && down && left && right
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 0 && dfs(row, col) {
				noOfClosedIsland++
			}
		}
	}

	return noOfClosedIsland
}
