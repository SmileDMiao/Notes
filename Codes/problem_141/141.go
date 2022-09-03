package leetcode

import "LeetCode/structures"

type ListNode = structures.ListNode

func hasCycle(head *ListNode) bool {

	cycle := false
	result := make(map[*ListNode]bool)

	for head != nil {
		result[head] = true

		if head.Next != nil && result[head.Next] {
			cycle = true
			break
		}

		head = head.Next
	}

	return cycle
}

func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}

	q := head
	p := head
	for p.Next != nil && p.Next.Next != nil && q.Next != nil {
		p = p.Next.Next
		q = q.Next

		if p == q {
			return true
		}
	}

	return false
}
