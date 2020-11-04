package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies), len(candies))
	curMax := max(candies)
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= curMax {
			res[i] = true
		}
	}
	return res
}

func max(arr []int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}
	res := arr[0]
	for i := 1; i < n; i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	fmt.Printf("输入: %v\n", candies)
	extraCandies := 3
	r := kidsWithCandies(candies, extraCandies)
	fmt.Printf("输出: %v\n", r)
}
