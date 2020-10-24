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
	str string
}

type res struct {
	correct_res bool
}

func Test_cases(t *testing.T) {
	test_cases := []test{
		{
			para{"()[]{}"},
			res{true},
		},

		{
			para{"{[]}"},
			res{true},
		},

		{
			para{"([)]"},
			res{false},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode 有效的括号------------------------\n")

	for _, p := range test_cases {
		cur_para, correct_ans := p.para.str, p.res.correct_res
		cur_res := isValid(cur_para)
		fmt.Printf("输入: %v, 输出:%v , 预期结果: %v\n", cur_para, cur_res, correct_ans)
	}
	fmt.Printf("\n\n\n")
}
