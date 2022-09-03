package leetcode

import "LeetCode/structures"

type ListNode = structures.ListNode

func reverseList(head *ListNode) *ListNode {
	var new *ListNode

	for head != nil {
		// 对之前的链表做头删
		node := head
		head = head.Next

		// 对新链表做头插
		node.Next = new
		new = node
	}
	return new
}
