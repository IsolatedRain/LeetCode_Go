package main

import "fmt"

func hasPath(maze [][]int, start []int, destination []int) bool {
	row := len(maze)
	col := len(maze[0])
	dir4 := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	visited := map[int]bool{}
	visited[start[0]*col+start[1]] = true
	q := [][]int{start}
	for len(q) > 0 {
		curPos := q[0]
		if curPos[0] == destination[0] && curPos[1] == destination[1] {
			return true
		}
		for _, d := range dir4 {
			nxtX := curPos[0] + d[0]
			nxty := curPos[1] + d[1]
			for 0 <= nxtX && nxtX < row && 0 <= nxty && nxty < col && maze[nxtX][nxty] == 0 {
				nxtX += d[0]
				nxty += d[1]
			}
			nxtX -= d[0]
			nxty -= d[1]
			curK := nxtX*col + nxty
			if !visited[curK] {
				q = append(q, []int{nxtX, nxty})
				visited[curK] = true
			}
		}
		q = q[1:]
	}
	return false
}

func main() {
	maze := [][]int{{0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}, {1, 1, 0, 1, 1}, {0, 0, 0, 0, 0}}
	start := []int{0, 4}
	destination := []int{1, 2}
	r := hasPath(maze, start, destination)
	fmt.Println(r)
}
