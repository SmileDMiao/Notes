// 817. Linked List Components
// question
// 给定链表头结点 head，该链表上的每个结点都有一个 唯一的整型值 。同时给定列表 G，该列表是上述链表中整型值的一个子集。返回列表 G 中组件的个数
// 这里对组件的定义为：链表中一段最长连续结点的值（该值必须在列表 G 中）构成的集合。
// example
// Input: head: 0->1->2->3, G = [0, 1, 3]; Output: 2; 链表中,0 和 1 是相连接的，且 G 中不包含 2，所以 [0, 1] 是 G 的一个组件，同理 [3] 也是一个组件，故返回 2。

// TODO
// 思路
// G转成map方便判断链表的Val是否在G中，遍历链表挨个判断是否在G中，连续出现算一次
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	m := make(map[int]struct{})
	for i := 0; i < len(G); i++ {
		m[G[i]] = struct{}{}
	}

	var result int
	var previous bool = false

	for head != nil {
		if _, ok := m[head.Val]; ok {
			if !previous {
				previous = true
				result++
			}
		} else {
			previous = false
		}

		head = head.Next
	}

	fmt.Println(result)
	return result
}

func createNode(v int) *ListNode {
	return &ListNode{Val: v}
}

func main() {
	G := []int{0, 1, 3, 4, 7, 9, 10}
	head := createNode(0)
	head.Next = createNode(1)
	head.Next.Next = createNode(2)
	head.Next.Next.Next = createNode(3)
	head.Next.Next.Next.Next = createNode(4)
	head.Next.Next.Next.Next.Next = createNode(6)
	head.Next.Next.Next.Next.Next.Next = createNode(7)
	head.Next.Next.Next.Next.Next.Next.Next = createNode(9)
	head.Next.Next.Next.Next.Next.Next.Next.Next = createNode(10)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = createNode(10)

	numComponents(head, G)
}
