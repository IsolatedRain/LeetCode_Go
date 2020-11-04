package main

import (
	"fmt"
)

func longestIncreasingPath(matrix [][]int) int {
	row := len(matrix)
	if row == 0 {
		return 0
	}

	col := len(matrix[0])
	outDegree := make([][]int, row, row)
	for i := 0; i < row; i++ {
		outDegree[i] = make([]int, col, col)
	}
	dir4 := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// 记录所有出度为0 的点
	// 即相对四个方向最大的点
	q := [][]int{}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			for _, v := range dir4 {
				x := r + v[0]
				y := c + v[1]
				if 0 <= x && x < row && 0 <= y && y < col && matrix[x][y] > matrix[r][c] {
					outDegree[r][c]++
				}
			}
			if outDegree[r][c] == 0 {
				q = append(q, []int{r, c})
			}
		}
	}

	res := 0
	for len(q) > 0 {
		res++
		size := len(q)
		for i := 0; i < size; i++ {
			r, c := q[0][0], q[0][1]
			q = q[1:]
			for _, v := range dir4 {
				x := r + v[0]
				y := c + v[1]
				if 0 <= x && x < row && 0 <= y && y < col && matrix[x][y] < matrix[r][c] {
					outDegree[x][y]--
					if outDegree[x][y] == 0 {
						q = append(q, []int{x, y})
					}
				}
			}
		}
	}

	return res
}

func main() {
	matrix := [][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1}}
	fmt.Printf("输入: %v\n", matrix)
	r := longestIncreasingPath(matrix)
	fmt.Printf("输出: %v\n", r)

}
