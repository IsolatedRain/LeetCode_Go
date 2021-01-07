package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	keyNums := 1
	visited := make([]bool, n)
	visited[0] = true
	var dfs func(curKey int)
	dfs = func(curKey int) {
		for _, key := range rooms[curKey] {
			if !visited[key] {
				visited[key] = true
				keyNums++
				dfs(key)
			}
		}
	}
	dfs(0)
	return keyNums == n
}

func main() {
	rooms := [][]int{{1}, {2}, {3}, {}}
	fmt.Println(canVisitAllRooms(rooms))
}
