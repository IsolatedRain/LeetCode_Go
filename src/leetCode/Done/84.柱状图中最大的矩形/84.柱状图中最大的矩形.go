package main

import (
	"fmt"
	_ "sort"
)

// 单调栈
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
	heights := []int{2, 1, 5, 6, 2, 3}
	// heights := []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15}
	r := largestRectangleArea(heights)
	fmt.Println(r)

}

// // 二分
// func largestRectangleArea(heights []int) int {
// 	n := len(heights)
// 	idx := []int{}
// 	for i := range heights {
// 		idx = append(idx, i)
// 	}
// 	sort.Slice(idx, func(i, j int) bool {
// 		return heights[idx[i]] < heights[idx[j]]
// 	})
// 	h := []int{-1, n}
// 	res := 0
// 	for _, i := range idx {
// 		curI := sort.Search(len(h), func(x int) bool {
// 			// fmt.Println(i, x)
// 			return h[x] >= i
// 		})
// 		R := append([]int{}, h[curI:]...)
// 		h = append(append(h[:curI], i), R...)
// 		res = max(res, (h[curI+1]-h[curI-1]-1)*heights[i])
// 	}
// 	return res
// }
