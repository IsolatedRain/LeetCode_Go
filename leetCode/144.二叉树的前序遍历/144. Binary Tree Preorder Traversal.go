package leetcode

import (
	"structures"
)

type treeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思路: Morris前序遍历,
// 利用curRoot 的左子节点的 最右子节点 为空 -> curRoot
// 如果 最右子节点.Right == curRoot, 表明 curRoot的左子树已经遍历完,
// 继续遍历 curRoot.Left, 即 curRoot = curRoot.Left.
// 中-> 左->右
func preorderTraversal(root *treeNode) []int {

	var vals []int
	var curRoot, curChild *treeNode = root, nil
	for curRoot != nil {
		if curRoot.Left == nil {
			vals = append(vals, curRoot.Val)
			curRoot = curRoot.Right
		} else {
			curChild = curRoot.Left
			for curChild.Right != nil && curChild.Right != curRoot {
				curChild = curChild.Right
			}
			if curChild.Right == nil {
				vals = append(vals, curRoot.Val)
				curChild.Right = curRoot
				curRoot = curRoot.Left
			} else {
				curChild.Right = nil
				curRoot = curRoot.Right
			}
		}
	}
	return vals
}

// for curRoot != nil {
// 	curChild = curRoot.Left
// 	if curChild != nil {
// 		for curChild.Right != nil && curChild.Right != curRoot {
// 			curChild = curChild.Right
// 		}
// 		if curChild.Right == nil {
// 			vals = append(vals, curRoot.Val)
// 			curChild.Right = curRoot
// 			curRoot = curRoot.Left
// 			continue
// 		}
// 		curChild.Right = nil
// 	} else {
// 		vals = append(vals, curRoot.Val)
// 	}
// 	curRoot = curRoot.Right
// }
// 	return vals
// }
