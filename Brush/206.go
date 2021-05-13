// 206. Reverse Linked List
// question
// 反转链表
// example
// Input [1,2,3,4,5], Output: [5,4,3,2,1]

// 思路
// 遍历链表，对老链表做头删除操作，对新链表做头插入操作

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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
