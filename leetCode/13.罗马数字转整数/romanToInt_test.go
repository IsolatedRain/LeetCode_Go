package leetCode

import (
	"fmt"
	"testing"
)

type test struct {
	roman
	res
}

type roman struct {
	p string
}

type res struct {
	correctRes int
}

func Test_Case(t *testing.T) {

	testCases := []test{
		{
			roman{"III"},
			res{3},
		},

		{
			roman{"IV"},
			res{4},
		},

		{
			roman{"IX"},
			res{9},
		},

		{
			roman{"LVIII"},
			res{58},
		},

		{
			roman{"MCMXCIV"},
			res{1994},
		},

		{
			roman{"MCMXICIVI"},
			res{2014},
		},
	}

	fmt.Printf("------------------------罗马数字 转 整数------------------------\n")

	for _, para := range testCases {
		correct_res, cur_roman := para.res.correctRes, para.roman.p
		cur_res := romanToInt(cur_roman)
		fmt.Printf("输入: %v    输出: %v  预期结果: %v\n", cur_roman, cur_res, correct_res)
	}
	fmt.Printf("\n\n\n")
}
