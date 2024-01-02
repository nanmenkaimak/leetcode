package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head

	for n > 0 && fast != nil {
		fast = fast.Next
		n--
	}

	if fast == nil {
		head = head.Next
		return head
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}
