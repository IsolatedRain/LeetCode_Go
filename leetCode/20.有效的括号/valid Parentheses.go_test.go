package leetcode

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
	correctRes bool
}

func Test_cases(t *testing.T) {
	testCases := []test{
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
	}

	fmt.Printf("------------------------Leetcode 有效的括号------------------------\n")

	for _, p := range testCases {
		curPara, correctAns := p.para.str, p.res.correctRes
		curRes := isValid(curPara)
		fmt.Printf("输入: %v, 输出:%v , 预期结果: %v\n", curPara, curRes, correctAns)
	}
	fmt.Printf("\n\n\n")
}
