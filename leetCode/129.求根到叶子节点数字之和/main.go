package main

import (
	"fmt"
	"structures"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode 自定义树节点 私有化
type TreeNode = structures.TreeNode
type pair struct {
	node *TreeNode
	v    int
}

// BFS
func sumNumbers(root *TreeNode) (sum int) {
	if root == nil {
		return
	}
	q := []pair{{root, root.Val}}
	for len(q) > 0 {
		curNode, curSum := q[0].node, q[0].v
		q = q[1:]
		if curNode.Left == nil && curNode.Right == nil {
			sum += curSum
		} else {
			if curNode.Left != nil {
				q = append(q, pair{curNode.Left, curSum*10 + curNode.Left.Val})
			}
			if curNode.Right != nil {
				q = append(q, pair{curNode.Right, curSum*10 + curNode.Right.Val})
			}
		}
	}
	return
}

// return getSum(root, 0) 递归
func getSum(r *TreeNode, curSum int) int {
	if r == nil {
		return 0
	}
	curSum = curSum*10 + r.Val
	if r.Left == nil && r.Right == nil {
		return curSum
	}
	return getSum(r.Left, curSum) + getSum(r.Right, curSum)
}

type test struct {
	para []int
	res  int
}

func main() {
	testCases := []test{
		{
			[]int{4, 9, 0, 5, 1},
			1026,
		},
		{
			[]int{1, 2, 3},
			25,
		},
		{
			[]int{},
			0,
		},
	}
	for _, p := range testCases {
		curPara := p.para
		fmt.Printf("%v\n", curPara)
		correctRes := p.res
		root := structures.Ints2TreeNode(curPara)
		curRes := sumNumbers(root)
		fmt.Printf("输入:%v\n输出:%v\n预期:%v\n", curPara, curRes, correctRes)
	}

}
