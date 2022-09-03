package leetcode

import (
	"LeetCode/structures"
	"fmt"
)

type ListNode structures.ListNode

func numComponents(head *structures.ListNode, G []int) int {
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
