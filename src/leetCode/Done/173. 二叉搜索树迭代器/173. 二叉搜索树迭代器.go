package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	cur *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{cur: root}
}

func (b *BSTIterator) Next() int {
	for {
		morrisRight := b.cur.Left
		if morrisRight != nil {
			for morrisRight.Right != nil && morrisRight != b.cur {
				morrisRight = morrisRight.Right
			}
			if morrisRight.Right == nil {
				morrisRight.Right = b.cur
				b.cur = b.cur.Left
				continue
			}
		}
		break
	}
	v := b.cur.Val
	b.cur = b.cur.Right
	return v
}

func (b *BSTIterator) HasNext() bool {
	return b.cur != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
