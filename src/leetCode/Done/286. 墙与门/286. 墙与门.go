package main

import (
	"math"
)

func wallsAndGates(rooms [][]int) {
	inf := math.MaxInt32
	dir4 := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	q := [][]int{}
	row, col := len(rooms), len(rooms[0])
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if rooms[r][c] == 0 {
				q = append(q, []int{r, c, 0})
			}
		}
	}
	for len(q) > 0 {
		curLen := len(q)
		for i := range q {
			r, c, curDistance := q[i][0], q[i][1], q[i][2]
			for _, m := range dir4 {
				newR, newC := r+m[0], c+m[1]
				if 0 <= newR && newR < row && 0 <= newC && newC < col &&
					rooms[newR][newC] == inf {
					q = append(q, []int{newR, newC, curDistance + 1})
					rooms[newR][newC] = curDistance + 1
				}
			}
		}
		q = q[curLen:]
	}
	return
}

func main() {
	rooms := [][]int{
		{2147483647, -1, 0, 2147483647},
		{2147483647, 2147483647, 2147483647, -1},
		{2147483647, -1, 2147483647, -1},
		{0, -1, 2147483647, 2147483647}}
	wallsAndGates(rooms)
}
