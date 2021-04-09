package main

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n, cur := 1, head
	for cur.Next != nil {
		cur = cur.Next
		n++
	}
	tail := cur
	mod := (n - k) % n
	cur = head
	for mod > 1 {
		cur = cur.Next
		mod--
	}
	h := cur.Next
	cur.Next = nil
	tail.Next = head
	return h
}
