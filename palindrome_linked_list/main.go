package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	h := head
	list1 := findMid(h)
	list2 := reverse(list1)

	for head != nil && list2 != nil {
		if head.Val != list2.Val {
			return false
		}
		head, list2 = head.Next, list2.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev, next *ListNode
	current := head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
