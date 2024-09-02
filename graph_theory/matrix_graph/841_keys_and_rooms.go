package matrix_graph

func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)
	canVisitAllRoomsDfs(0, rooms, visited)

	return len(visited) == len(rooms)
}

func canVisitAllRoomsDfs(current int, rooms [][]int, visited map[int]bool) {
	visited[current] = true
	for _, key := range rooms[current] {
		if !visited[key] {
			canVisitAllRoomsDfs(key, rooms, visited)
		}
	}
}
