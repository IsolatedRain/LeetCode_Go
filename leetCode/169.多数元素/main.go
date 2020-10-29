package main

import "fmt"

// 摩尔投票法, 逐个抵消.
func majorityElement(nums []int) int {
	cnt := 1
	res := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] == res {
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				res = nums[i+1] // 因为必定存在, 所以不会越界
			}
		}
	}
	return res
}

func main() {
	p := []int{2, 2, 1, 1, 1, 2, 2}
	r := majorityElement(p)
	fmt.Printf("%v\n", r)

}
