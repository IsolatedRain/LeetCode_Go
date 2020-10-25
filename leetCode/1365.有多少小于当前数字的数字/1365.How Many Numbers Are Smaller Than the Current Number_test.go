package leetCode

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
		cur_para := p.nums
		correct_res := p.res
		cur_res := smallerNumbersThanCurrent(cur_para)
		fmt.Printf("输入:%v,    输出:%v,     预期结果:%v    \n", cur_para, cur_res, correct_res)
		same := true
		for i, v := range cur_res {
			if correct_res[i] != v {
				same = false
				break
			}
		}
		fmt.Printf("结果相同:%v\n", same)
	}
}
