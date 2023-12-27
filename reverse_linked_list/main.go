package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var p, c, n *ListNode
	p = nil
	c = head

	for c != nil {
		n = c.Next
		c.Next = p
		p = c
		c = n
	}
	head = p
	return head
}
