package leetcode

import (
	"fmt"
	"testing"
)

func createNode(value int, nextNode *ListNode) *ListNode {
	return &ListNode{Val: value, Next: nextNode}
}

func Test_Problem136(t *testing.T) {

	root := createNode(3, nil)
	root.Next = createNode(2, createNode(0, createNode(4, root)))

	fmt.Println(hasCycle(root))
	fmt.Println(hasCycle2(root))

}
