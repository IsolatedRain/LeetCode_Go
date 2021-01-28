package main

import (
	"fmt"
	"math"
)

func solve(nums []int, queries [][]int) []int {
	n := len(nums)
	// mod := int(math.Pow10(9)) + 7
	mod := 1000000007
	// sqrt := int(math.Pow(float64(n), 0.5)) + 1
	sqrt := int(math.Sqrt(float64(n))) + 1
	d := make([][]int, sqrt)
	for i := 1; i < sqrt; i++ {
		d[i] = make([]int, n)
		for j := n - 1; j >= n-i; j-- {
			d[i][j] = nums[j]
		}
		for j := n - i - 1; j >= 0; j-- {
			d[i][j] = (nums[j] + d[i][j+i]) % mod
		}
	}

	res := []int{}
	for _, q := range queries {
		x, y := q[0], q[1]
		if y > sqrt-1 {
			cur := 0
			for i := x; i < n; i += y {
				cur += nums[i]
				cur %= mod
			}
			res = append(res, cur)
		} else {
			res = append(res, d[y][x])
		}
	}
	return res
}

func main() {
	fmt.Println(solve([]int{0, 1, 2, 3, 4, 5, 6, 7}, [][]int{{0, 3}, {5, 1}, {4, 2}}))
}
