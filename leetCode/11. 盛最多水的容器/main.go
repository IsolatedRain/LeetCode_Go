package main

import "fmt"

// 左右指针,逐步缩进
// 容量 = width * height
// width 固定, 移动 较小的 height 才可能增大容量.
func maxArea(height []int) int {
	L, R := 0, len(height)-1
	res := 0
	for L < R {
		curWidth := R - L
		if height[L] <= height[R] {
			res = max(res, curWidth*height[L])
			L++
		} else {
			res = max(res, curWidth*height[R])
			R--
		}
	}
	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	p := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	r := maxArea(p)
	fmt.Printf("输入:%v\n输出:%v", p, r)
}
