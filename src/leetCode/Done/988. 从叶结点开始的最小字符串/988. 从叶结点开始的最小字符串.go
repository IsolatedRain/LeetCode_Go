package main

// TreeNode 树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	res := ""
	var dfs func(r *TreeNode, path []byte)
	dfs = func(r *TreeNode, path []byte) {
		if r != nil {
			if r.Left == nil && r.Right == nil {
				res = min(res, string(path))
			}
			if r.Left != nil {
				dfs(r.Left, append(path, byte(r.Left.Val)+'a'))
			}
			if r.Right != nil {
				dfs(r.Right, append(path, byte(r.Right.Val)+'a'))
			}
		}
	}
	dfs(root, []byte{byte(root.Val) + 'a'})
	return res
}
func min(x, y string) string {
	xLen, yLen := len(x), len(y)
	if xLen == 0 {
		return y
	}
	if yLen == 0 {
		return x
	}
	i := 0
	for i < xLen && i < yLen {
		if x[i] > y[i] {
			return y
		} else if x[i] < y[i] {
			return x
		}
	}
	if xLen > yLen {
		return y
	}
	return x
}
