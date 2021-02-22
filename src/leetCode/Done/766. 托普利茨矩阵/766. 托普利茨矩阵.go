package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	d := map[int]int{}
	row, col := len(matrix), len(matrix[0])
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			k := r - c
			if _, ok := d[k]; !ok {
				d[k] = matrix[r][c]
			} else {
				if d[k] != matrix[r][c] {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	m := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	fmt.Println(isToeplitzMatrix(m))

}
