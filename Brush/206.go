package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var new *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = new
		new = head
		head = tmp
	}
	return new
}
