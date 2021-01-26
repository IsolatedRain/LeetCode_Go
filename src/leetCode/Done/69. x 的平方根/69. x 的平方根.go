package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	o, res := float64(x), float64(x)
	for {
		cur := 0.5 * (res + o/res)
		if abs(res-cur) < 1e-7 {
			break
		}
		res = cur
	}
	return int(res)
}
func abs(x float64) float64 {
	if x < 0 {
		return x * -1
	}
	return x
}
func main() {
	fmt.Println(mySqrt(20))
}
