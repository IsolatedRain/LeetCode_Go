package main

import (
	"fmt"
)

// 深搜 回溯, mark数组标记 已选位置
func permute(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0, n)
	mark := make([]int, n)
	var dfs func(curList []int)
	dfs = func(curList []int) {
		if len(curList) == n {
			res = append(res, curList)
			return
		}
		for i := 0; i < n; i++ {
			if mark[i] == 0 {
				mark[i] = 1 // 标记
				dfs(append(curList, nums[i]))
				mark[i] = 0 // 移除标记
			}
		}
	}
	dfs([]int{})
	return res
}

func main() {
	para := []int{1, 2, 3, 4, 5}
	r := permute(para)
	fmt.Printf("输入: %v\n输出: %v\n", para, r)
}
