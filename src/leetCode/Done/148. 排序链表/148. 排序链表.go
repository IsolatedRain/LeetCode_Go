package main

import "fmt"

// ListNode 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	right := partition(head)
	left := head
	L := sortList(left)
	R := sortList(right)
	return merge(L, R)
}
func partition(h *ListNode) *ListNode {
	fast, slow := h, h
	tail := new(ListNode)
	tail.Next = slow
	for fast != nil && fast.Next != nil {
		tail = tail.Next
		slow = slow.Next
		fast = fast.Next.Next
	}
	tail.Next = nil
	return slow
}
func merge(L1, L2 *ListNode) *ListNode {
	p1, p2 := L1, L2
	r := new(ListNode)
	cur := r
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			cur.Next = p2
			p2 = p2.Next
		} else {
			cur.Next = p1
			p1 = p1.Next
		}
		cur = cur.Next
	}
	if p1 != nil {
		cur.Next = p1
	} else {
		cur.Next = p2
	}
	return r.Next
}

func main() {
	vals := []int{4, 2, 1, 3}
	head := &ListNode{Val: vals[0]}
	cur := head
	for _, v := range vals[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	fmt.Println(sortList(head))
}
