package main

import "fmt"

func maxScore(cardPoints []int, k int) int {
	sum, n := 0, len(cardPoints)
	for i := range cardPoints {
		sum += cardPoints[i]
	}
	if n <= k {
		return sum
	}
	curSum, L, R := 0, 0, n-k
	for i := L; i < R; i++ {
		curSum += cardPoints[i]
	}
	res := sum - curSum
	for R < n {
		curSum -= cardPoints[L]
		curSum += cardPoints[R]
		res = max(res, sum-curSum)
		L++
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
	// cardPoints := []int{9, 7, 7, 9, 7, 7, 9}
	// k := 4
	cardPoints := []int{96, 90, 41, 82, 39, 74, 64, 50, 30}
	k := 8
	fmt.Println(maxScore(cardPoints, k))
}

// 左右前缀和
// func maxScore(cardPoints []int, k int) int {
// 	n := len(cardPoints)
// 	left := []int{0}
// 	right := []int{0}
// 	for i := range cardPoints {
// 		left = append(left, left[len(left)-1]+cardPoints[i])
// 		right = append(right, right[len(right)-1]+cardPoints[n-i-1])
// 	}
// 	if n <= k {
// 		return left[len(left)-1]
// 	}
// 	res := 0
// 	for i := 0; i <= k; i++ {
// 		res = max(res, left[i]+right[k-i])
// 	}
// 	return res
// }
