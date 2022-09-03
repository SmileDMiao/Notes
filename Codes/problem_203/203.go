package leetcode

import "LeetCode/structures"

type ListNode = structures.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head

	tmp := dummy
	for tmp != nil && tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return dummy.Next
}
