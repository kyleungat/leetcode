package matrix_graph

func maxAreaOfIsland(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	var maxArea int

	var search func(row, col, counter int) int
	search = func(row, col, counter int) int {
		if row < 0 || col < 0 || row > rows-1 || col > cols-1 || grid[row][col] == 0 {
			return 0
		}
		grid[row][col] = 0

		leftArea := search(row, col-1, counter)
		rightArea := search(row, col+1, counter)
		topArea := search(row-1, col, counter)
		downArea := search(row+1, col, counter)

		counter = counter + 1 + leftArea + rightArea + topArea + downArea
		return counter
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 {
				var counter int
				area := search(row, col, counter)
				if area > maxArea {
					maxArea = area
				}
			}

		}
	}

	return maxArea
}
