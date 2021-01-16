package main

import "fmt"

func majorityElement(nums []int) int {
	curNum := nums[0]
	curCnt := 1
	for _, num := range nums[1:] {
		if curCnt == 0 {
			curNum = num
			curCnt = 1
			continue
		}
		if num == curNum {
			curCnt++
		} else {
			curCnt--
		}
	}
	return curNum
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2, 1, 1, 3, 3, 3, 3, 3, 3, 3, 1, 1, 1, 2, 3, 3, 4, 2, 1, 4, 2, 1, 3, 4, 2, 321, 1, 2, 1, 1, 1, 2, 2, 3, 4, 2, 1}
	fmt.Println(majorityElement(nums))
}
