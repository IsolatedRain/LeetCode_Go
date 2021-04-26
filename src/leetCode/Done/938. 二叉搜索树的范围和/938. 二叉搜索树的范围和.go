// https://leetcode-cn.com/problems/range-sum-of-bst/
package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var dfs func(r *TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		if r.Val > high {
			return dfs(r.Left)
		}
		if r.Val < low {
			return dfs(r.Right)
		}
		return r.Val + dfs(r.Left) + dfs(r.Right)
	}
	return dfs(root)
}
