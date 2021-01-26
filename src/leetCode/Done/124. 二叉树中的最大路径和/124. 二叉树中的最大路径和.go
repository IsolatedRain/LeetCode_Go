package main

import "fmt"

//TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := root.Val
	var dfs func(r *TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		if r.Left == nil && r.Right == nil {
			return r.Val
		}
		left := dfs(r.Left)
		right := dfs(r.Right)
		curMax := max(r.Val, r.Val+left, r.Val+right, r.Val+left+right)
		res = max(res, curMax)
		return res
	}
	dfs(root)
	return res
}
func max(x ...int) int {
	ret := x[0]
	for _, v := range x[1:] {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func main() {
	vals := []int{1, 2, 3}
	r := &TreeNode{Val: vals[0]}
	r.Left = &TreeNode{Val: vals[1]}
	r.Right = &TreeNode{Val: vals[2]}
	fmt.Println(maxPathSum(r))
}
