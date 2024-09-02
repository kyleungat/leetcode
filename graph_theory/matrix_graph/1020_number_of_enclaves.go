package matrix_graph

func numEnclaves(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	var cells int

	var dfs func(row, col, counter int) (bool, int)
	dfs = func(row, col, counter int) (bool, int) {
		if (row == 0 || col == 0 || row == rows-1 || col == cols-1) && grid[row][col] == 1 { // boundary
			return true, 0
		}

		if grid[row][col] != 1 {
			return false, 0
		}

		grid[row][col] = 0
		left, walkedLeft := dfs(row, col-1, counter)
		right, walkedRight := dfs(row, col+1, counter)
		top, walkedTop := dfs(row-1, col, counter)
		down, walkedDown := dfs(row+1, col, counter)
		isBoundary := left || right || top || down
		counter = counter + 1 + walkedLeft + walkedRight + walkedTop + walkedDown

		return isBoundary, counter
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 {
				var counter int
				isBoundary, walked := dfs(row, col, counter)
				if !isBoundary {
					cells += walked
				}
			}
		}
	}
	return cells
}
