package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func createNode(val int) *ListNode {
	return &ListNode{val, nil}
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		if head.Val == val {
			tmp := head.Next
			head.Next = nil
			head = tmp
		}
	}

	tmp := head
	for tmp != nil && tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return head
}

func removeElements1(head *ListNode, val int) *ListNode {
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
