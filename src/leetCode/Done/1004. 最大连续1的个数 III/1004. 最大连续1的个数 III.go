package main

import "fmt"

func longestOnes(A []int, K int) int {
	c := 0
	for _, v := range A {
		if v == 0 {
			c++
		}
	}
	if c <= K {
		return len(A)
	}

	A = append(A, 0)
	L, R, res, count0, n := 0, 0, 0, 0, len(A)
	for R < n {
		if A[R] == 0 {
			count0++
			if count0 > K {
				res = max(res, R-L)
				for A[L] != 0 {
					L++
				}
				L++
				count0--
			}
		}
		R++
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
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}
