package main

import (
	"fmt"
	"math"
)

func smallestFactorization(a int) int {
	if a <= 9 {
		return a
	}
	inf := math.MaxInt32
	res := 0
	power := 1
	for a > 1 {
		for i := 9; i >= 1; i-- {
			if a%i == 0 {
				res += i * power
				a /= i
				break
			}
		}
		if res > inf {
			return 0
		}
		power *= 10
	}
	return res
}

func main() {
	fmt.Println(smallestFactorization(18000000))
}
