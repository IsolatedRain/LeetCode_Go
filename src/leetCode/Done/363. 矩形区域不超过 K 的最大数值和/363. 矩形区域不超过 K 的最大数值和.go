package main

import (
	"fmt"
	"math"
)

func maxSumSubmatrix(matrix [][]int, k int) int {
	n, m := len(matrix), len(matrix[0])
	res := math.MinInt64
	for i := range matrix { // 上边界
		curColSum := make([]int, m)
		for j := i; j < n; j++ { // 下边界
			for c := 0; c < m; c++ { // 当前列合
				if matrix[j][c] == k {
					return k
				}
				curColSum[c] += matrix[j][c]
				if curColSum[c] == k {
					return k
				}
			}
			if v := arrMaxSubSum(k, curColSum...); v <= k {
				res = max(res, v)
				fmt.Println(v)
				continue
			}
			for L := 0; L < m; L++ {
				curSum := 0
				for R := L; R < m; R++ {
					curSum += curColSum[R]
					if curSum == k {
						return k
					}
					if curSum <= k {
						res = max(res, curSum)
					}
				}
			}
		}
	}
	return res
}
func arrMaxSubSum(k int, arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res = max(res, max(arr[i], arr[i]+arr[i-1]))
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
	// matrix := [][]int{{1, 0, 1}, {0, -2, 3}}
	matrix := [][]int{{2, 2, -1}}
	k := 0
	fmt.Println(maxSumSubmatrix(matrix, k))
}
