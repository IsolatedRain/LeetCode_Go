package leetCode

import (
	"fmt"
	"testing"
)

type test struct {
	arr []int
	res int
}

func Test_cases(t *testing.T) {
	testCases := []test{
		{
			[]int{2, 1, 4, 7, 3, 2, 5},
			5,
		},
		{
			[]int{2, 2, 2},
			0,
		},
		{
			[]int{},
			0,
		},
	}
	fmt.Println("------------------------数组中的最长山脉---------------------------")
	for _, p := range testCases {
		cur_para := p.arr
		correct_res := p.res
		cur_res := longestMountain(cur_para)
		fmt.Printf("输入:%v,    输出:%v,     预期结果:%v    \n", cur_para, cur_res, correct_res)
		fmt.Printf("结果正确:  %v\n", correct_res == cur_res)
	}
}
