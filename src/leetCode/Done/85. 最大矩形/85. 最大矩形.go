package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	row := len(matrix)
	col := len(matrix[0])
	heights := make([][]int, row)
	for r := range heights {
		heights[r] = make([]int, col)
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if matrix[r][c] == '1' {
				if r == 0 {
					heights[r][c] = 1
				} else {
					heights[r][c] = heights[r-1][c] + 1
				}
			}
		}
	}
	res := 0
	for _, height := range heights {
		res = max(res, largestRectangleArea(height))
	}
	return res
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := []int{-1}
	res := 0
	for i, h := range heights {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > h {
			curH := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			res = max(res, (i-stack[len(stack)-1]-1)*curH)
		}
		stack = append(stack, i)
	}
	for len(stack) > 1 {
		curH := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		res = max(res, (n-stack[len(stack)-1]-1)*curH)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	r := maximalRectangle(matrix)
	fmt.Println(r)
}
