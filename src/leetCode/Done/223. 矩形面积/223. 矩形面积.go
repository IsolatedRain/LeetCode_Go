package main

import "fmt"

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	long := max(0, min(C, G)-max(A, E))
	height := max(0, min(D, H)-max(B, F))
	areaA := (C - A) * (D - B)
	areaB := (G - E) * (H - F)
	return areaA + areaB - long*height
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
}
