package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x /= 10
		if !isValid(res) {
			return 0
		}
	}
	return res
}
func isValid(x int) bool {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return false
	}
	return true
}

func main() {
	fmt.Println(reverse(-123))
}
