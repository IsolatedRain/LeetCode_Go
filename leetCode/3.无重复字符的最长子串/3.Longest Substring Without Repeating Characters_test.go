package leetcode

import (
	"fmt"
	"testing"
)

type test struct {
	s   string
	res int
}

func Test_Case(t *testing.T) {

	testCases := []test{
		{
			"abcabcbb",
			3,
		},

		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
		{
			"",
			0,
		},
	}
	fmt.Printf("------------------------无重复字符的最长子串------------------------\n")

	for _, curCase := range testCases {
		correctRes, curS := curCase.res, curCase.s
		curRes := lengthOfLongestSubstring(curS)
		fmt.Printf("输入: %v,    输出: %v,  预期结果: %v\n", curS, curRes, correctRes)
		fmt.Printf("%v\n", curRes == correctRes)
	}
	fmt.Printf("\n\n\n")
}
