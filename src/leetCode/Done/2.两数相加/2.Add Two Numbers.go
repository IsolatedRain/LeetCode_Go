package leetcode

import "structures"

// "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  */
type listNode = structures.ListNode

func addTwoNumbers(l1 *listNode, l2 *listNode) *listNode {
	head := &listNode{}
	cur := head
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		curVal := carry % 10
		carry /= 10
		cur.Next = &listNode{Val: curVal}
		cur = cur.Next
	}
	return head.Next
}
