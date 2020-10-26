package leetcode

import (
	"fmt"
	"testing"
)

type test struct {
	arr []int
	tar int
	res []int
}

func Test_cases(t *testing.T) {

	test_cases := []test{
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},

		{
			[]int{3, 2, 4},
			5,
			[]int{0, 1},
		},

		{
			[]int{0, 8, 7, 3, 3, 4, 2},
			11,
			[]int{1, 3},
		},

		{
			[]int{0, 1},
			1,
			[]int{0, 1},
		},

		{
			[]int{0, 3},
			5,
			[]int{},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode 两数之和------------------------\n")

	for _, p := range test_cases {
		correct_ans, cur_arr, cur_tar := p.res, p.arr, p.tar
		cur_res := twoSum(cur_arr, cur_tar)
		fmt.Printf("输入:%v, target:%v      输出:%v   预期结果: %v  \n", cur_arr, cur_tar, cur_res, correct_ans)
		len_correct_ans := len(correct_ans)
		len_cur_res := len(cur_res)
		same := true
		if len_correct_ans != len_cur_res {
			same = false
		} else {
			for i, v := range correct_ans {
				if cur_res[i] != v {
					same = false
					break
				}
			}
		}

		fmt.Printf("结果: %v\n", same)
	}
	fmt.Printf("\n\n\n")
}
