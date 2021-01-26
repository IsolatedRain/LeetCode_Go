package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Solution 存数
type Solution struct {
	arr []int
}

// Constructor 初始化
func Constructor(nums []int) Solution {
	return Solution{arr: nums}
}

// Pick 随机选择
func (s *Solution) Pick(target int) int {
	rand.Seed(time.Now().UnixNano())
	curCount := 0
	res := -1
	for i := range s.arr {
		if s.arr[i] == target {
			curCount++
			if rand.Intn(curCount) == 0 {
				res = i
			}
		}
	}
	return res
}

func main() {
	obj := Constructor([]int{1, 2, 3, 3, 3})
	fmt.Println(obj.Pick(3))
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
