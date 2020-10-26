package leetcode

import (
	"fmt"
	"testing"
)

type test struct {
	nums []int
	res  []int
}

func Test_cases(t *testing.T) {
	testCases := []test{
		{
			[]int{8, 1, 2, 2, 3},
			[]int{4, 0, 1, 1, 3},
		},
		{
			[]int{6, 5, 4, 8},
			[]int{2, 1, 0, 3},
		},
		{
			[]int{7, 7, 7, 7},
			[]int{0, 0, 0, 0},
		},
	}
	fmt.Println("------------------------有多少小于当前数字的数字---------------------------")
	for _, p := range testCases {
		curPara := p.nums
		correctRes := p.res
		curRes := smallerNumbersThanCurrent(curPara)
		fmt.Printf("输入:%v,    输出:%v,     预期结果:%v    \n", curPara, curRes, correctRes)
		same := true
		for i, v := range curRes {
			if correctRes[i] != v {
				same = false
				break
			}
		}
		fmt.Printf("结果相同:%v\n", same)
	}
}
