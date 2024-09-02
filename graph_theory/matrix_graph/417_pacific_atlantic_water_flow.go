package matrix_graph

// invalid case
// valid move
// base case
// other case
// visited

func pacificAtlantic(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])
	var result [][]int

	var isFlowToPacific func(row, col, previousHeight int, visited map[string]bool) bool
	isFlowToPacific = func(row, col, previousHeight int, visited map[string]bool) bool {
		if row < 0 || col < 0 || row > rows-1 || col > cols-1 {
			return false
		}
		key := string(byte(row)) + string(byte(col))
		if heights[row][col] <= previousHeight && !visited[key] {
			visited[key] = true
			if row == 0 || col == 0 { // top edge and left edge
				return true
			}
			return isFlowToPacific(row-1, col, heights[row][col], visited) ||
				isFlowToPacific(row, col-1, heights[row][col], visited) ||
				isFlowToPacific(row, col+1, heights[row][col], visited) ||
				isFlowToPacific(row+1, col, heights[row][col], visited)
		}
		return false
	}

	var isFlowToAtlantic func(row, col, previousHeight int, visited map[string]bool) bool
	isFlowToAtlantic = func(row, col, previousHeight int, visited map[string]bool) bool {
		if row < 0 || col < 0 || row > rows-1 || col > cols-1 {
			return false
		}
		key := string(byte(row)) + string(byte(col))
		if heights[row][col] <= previousHeight && !visited[key] {
			visited[key] = true
			if row == rows-1 || col == cols-1 { // top edge and left edge
				return true
			}
			return isFlowToAtlantic(row, col+1, heights[row][col], visited) ||
				isFlowToAtlantic(row+1, col, heights[row][col], visited) ||
				isFlowToAtlantic(row, col-1, heights[row][col], visited) ||
				isFlowToAtlantic(row-1, col, heights[row][col], visited)
		}
		return false
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if isFlowToPacific(row, col, heights[row][col], make(map[string]bool)) && isFlowToAtlantic(row, col, heights[row][col], make(map[string]bool)) {
				result = append(result, []int{row, col})
			}
		}
	}
	return result
}
