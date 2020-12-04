package main

import "fmt"

// 检查 v-1 结尾的有几个，能否加进去
// 如果不能， 检查 v+1 v+2 是否 >0， 即能否凑成新的子序列.
// 如果不能， 检查 v+1 v+2 是否 >0， 即能否凑成新的子序列.
func isPossible(nums []int) bool {
	count := map[int]int{}
	for _, v := range nums {
		count[v]++
	}
	preEndCount := map[int]int{}
	for _, v := range nums {
		if count[v] > 0 {
			if preCount, ok := preEndCount[v-1]; ok && preCount > 0 {
				preEndCount[v-1]--
				preEndCount[v]++
				count[v]--
			} else if count[v+1] > 0 && count[v+2] > 0 {
				preEndCount[v+2]++
				count[v]--
				count[v+1]--
				count[v+2]--
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	nums := []int{1, 2, 3, 3, 4, 4, 5, 5}
	r := isPossible(nums)
	fmt.Println(r)
}
