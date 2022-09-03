package leetcode

import "LeetCode/structures"

type ListNode = structures.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	count := 0
	first, second := head, head

	for first != nil {
		if count > n {
			second = second.Next
		}
		first = first.Next
		count += 1
	}

	if count == n {
		return head.Next
	}
	second.Next = second.Next.Next

	return head
}


func main() {

}
