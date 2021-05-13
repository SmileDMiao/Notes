// 203. Remove Linked List Elements
// question
// 删除链表中value为指定值的节点
// example
// Input [1,2,6,3,4,5,6], val = 6, Output: [1,2,3,4,5]

// 思路
// 添加一个空链表头，这样每个节点的操作就一致了

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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
