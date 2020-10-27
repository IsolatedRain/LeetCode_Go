package leetcode

import (
	"fmt"
	"structures"
	"testing"
)

type test struct {
	l1  []int
	l2  []int
	res []int
}

func Test_Cases(t *testing.T) {
	testCases := []test{
		{
			[]int{2, 4, 3},
			[]int{5, 6, 4},
			[]int{7, 0, 8},
		},
		{
			[]int{9, 9, 9, 9, 9},
			[]int{5},
			[]int{4, 0, 0, 0, 0, 1},
		},
		{
			[]int{9, 9, 9, 9, 9},
			[]int{1},
			[]int{0, 0, 0, 0, 0, 1},
		},
	}

	fmt.Printf("------------------------两数相加------------------------\n")

	for _, p := range testCases {
		curL1, curL2, correctRes := p.l1, p.l2, p.res
		L1 := structures.Ints2List(curL1)
		L2 := structures.Ints2List(curL2)
		listRes := addTwoNumbers(L1, L2)
		curRes := structures.List2Ints(listRes)
		fmt.Printf("输入:%v  %v,       输出:%v,    预期结果:%v\n", curL1, curL2, curRes, correctRes)
	}
	fmt.Printf("\n\n\n")
}
