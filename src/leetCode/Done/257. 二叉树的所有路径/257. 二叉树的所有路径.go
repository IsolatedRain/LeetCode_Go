// Definition for a binary tree node.
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	path := []int{}
	res := []string{}
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r.Left == nil && r.Right == nil {
			path = append(path, r.Val)
			t := ""
			for i := range path {
				t += fmt.Sprintf("%d", path[i]) + "->"
			}
			res = append(res, t[:len(t)-2])
			path = path[:len(path)-1]
			return
		}
		path = append(path, r.Val)
		if r.Left != nil {
			dfs(r.Left)
		}
		if r.Right != nil {
			dfs(r.Right)
		}
		path = path[:len(path)-1]
	}
	dfs(root)
	return res
}
