package main

import (
	"fmt"
	"structures"
	"testing"
)

type test struct {
	para []int
	res  []int
}

var null = structures.NULL

func Test_Problem144(t *testing.T) {

	testCases := []test{
		{
			[]int{1, null, 2, 3},
			[]int{1, 2, 3},
		},

		{
			[]int{1},
			[]int{1},
		},

		{
			[]int{},
			[]int{},
		},
	}

	fmt.Printf("------------------------二叉树的前序遍历------------------------\n\n\n")

	for _, p := range testCases {
		curPara, correctRes := p.para, p.res
		root := structures.Ints2TreeNode(curPara)
		curRes := preorderTraversal(root)
		fmt.Printf("输入:%v, 输出:%v    预期:%v  \n", curPara, curRes, correctRes)
	}
	fmt.Printf("\n\n\n")
}
