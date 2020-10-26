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
	strs []string
}

type res struct {
	correctRes string
}

func Test_Case(t *testing.T) {

	testCases := []test{
		{
			para{[]string{"flower", "flow", "flight"}},
			res{"fl"},
		},

		{
			para{[]string{"dog", "racecar", "car"}},
			res{""},
		},
	}

	fmt.Printf("------------------------最长公共前缀------------------------\n")

	for _, para := range testCases {
		correct_res, cur_strs := para.res.correctRes, para.strs
		cur_res := longestCommonPrefix(cur_strs)
		fmt.Printf("输入: %v,    输出: %v,  预期结果: %v\n", cur_strs, cur_res, correct_res)
	}
	fmt.Printf("\n\n\n")
}
