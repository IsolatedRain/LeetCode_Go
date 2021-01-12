package main

import (
	"fmt"
)

func paintingPlan(n int, k int) int {
	sum := n * n
	if sum == k || k == 0 {
		return 1
	}
	isValid := func(r int) bool {
		curK := k - r*n
		leftPerCol := n - r
		return curK%leftPerCol == 0
	}
	cal := func(r int) int {
		rowNums := fac(n) / (fac(n-r) * fac(r))
		m := (k - r*n) / (n - r)
		colNums := fac(n) / (fac(n-m) * fac(m))
		return rowNums * colNums
	}
	res := 0
	for r := 0; r < n; r++ {
		if isValid(r) {
			res += cal(r)
		}
	}
	return res
}
func fac(x int) int {
	r := 1
	for i := 1; i <= x; i++ {
		r *= i
	}
	return r
}

func main() {
	n := 6
	k := 24
	fmt.Println(paintingPlan(n, k))
}
