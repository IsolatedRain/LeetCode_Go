package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	p := dummy
	for i := 1; i < left; i++ {
		p = p.Next
	}
	h := p.Next
	node := p.Next
	for i := left; i < right; i++ {
		cur := node.Next
		tmp := cur.Next
		cur.Next = h
		node.Next = tmp
		h = cur
	}
	p.Next = h
	return dummy.Next
}
