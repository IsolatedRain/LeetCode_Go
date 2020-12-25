package main

import "fmt"

func numTeams(rating []int) int {
	n := len(rating)
	leftBigger := make([]int, n)
	leftLess := make([]int, n)
	rightBigger := make([]int, n)
	rightLess := make([]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if rating[j] > rating[i] {
				rightBigger[i]++
			} else if rating[j] < rating[i] {
				rightLess[i]++
			}
		}
		for j := 0; j < i; j++ {
			if rating[i] > rating[j] {
				leftLess[i]++
			} else if rating[i] < rating[j] {
				leftBigger[i]++
			}
		}
	}
	res := 0
	for i := range rating {
		res += leftLess[i] * rightBigger[i]
		res += leftBigger[i] * rightLess[i]
	}
	return res
}

func main() {
	rating := []int{2, 5, 3, 4, 1}
	r := numTeams(rating)
	fmt.Println(r)
}
