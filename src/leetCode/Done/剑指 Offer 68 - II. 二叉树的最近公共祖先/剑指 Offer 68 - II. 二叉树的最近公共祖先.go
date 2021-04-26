// https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/
package main

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(r *TreeNode) *TreeNode
	dfs = func(r *TreeNode) *TreeNode {
		if r == nil || r == p || r == q {
			return r
		}
		left := dfs(r.Left)
		right := dfs(r.Right)
		if left != nil && right != nil {
			return r
		}
		if left == nil {
			return right
		}
		if right == nil {
			return left
		}
		return nil
	}

	return dfs(root)
}

func main() {

}
