package leetCode

import (
	"fmt"
	"testing"
)

type test struct {
	para
	res
}

type para struct {
	nums   []int
	target int
}

type res struct {
	correct_res []int
}

func Test_cases(t *testing.T) {

	test_cases := []test{
		{
			para{[]int{3, 2, 4}, 6},
			res{[]int{1, 2}},
		},

		{
			para{[]int{3, 2, 4}, 5},
			res{[]int{0, 1}},
		},

		{
			para{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			res{[]int{1, 3}},
		},

		{
			para{[]int{0, 1}, 1},
			res{[]int{0, 1}},
		},

		{
			para{[]int{0, 3}, 5},
			res{[]int{}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode 两数之和------------------------\n")

	for _, p := range test_cases {
		correct_ans, cur_para := p.res.correct_res, p.para
		cur_res := twoSum(cur_para.nums, cur_para.target)
		fmt.Printf("输入:%v, target:%v      输出:%v   预期结果: %v  \n", cur_para.nums, cur_para.target, cur_res, correct_ans)
	}
	fmt.Printf("\n\n\n")
}
