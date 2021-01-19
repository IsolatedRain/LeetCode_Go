package main

import "fmt"

func missingNumber(nums []int) int {
	res := len(nums)
	for i := range nums {
		res ^= nums[i] ^ i
	}
	return res
	// n := len(nums)
	// i := 0
	// for i < n {
	// 	for nums[i] != i {
	// 		if nums[i] == n {
	// 			nums[i], nums[n-1] = nums[n-1], nums[i]
	// 		} else {
	// 			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	// 		}
	// 		if nums[i] == n {
	// 			break
	// 		}
	// 	}
	// 	i++
	// }
	// for i := range nums {
	// 	if nums[i] != i {
	// 		return i
	// 	}
	// }
	// return n
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums))
}
