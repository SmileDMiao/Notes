// 19. Remove Nth Node From End of List
// question
// 删除链表倒数第N个节点
// example
// Input: [2, 7, 11, 15], target = 3; Output [2,7,15]

// 思路(一次循环)
// 设 first second都是原链表, 遍历first, 当遍历到节点数 > n的时候，
// 开始移动second,first second两个指针的距离是n那么遍历结束，second就位于倒数第N个节点

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func createNode(val int) *ListNode {
	return &ListNode{Val: val}
}
func main() {
	node := createNode(1)
	node.Next = createNode(2)
	node.Next.Next = createNode(3)
	node.Next.Next.Next = createNode(4)
	node.Next.Next.Next.Next = createNode(5)

	removeNthFromEnd(node, 2)
}
