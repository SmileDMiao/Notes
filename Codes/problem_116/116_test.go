package leetcode

import (
	"fmt"
	"testing"
)

func createNode(value int) *node {
	return &node{value, nil, nil, nil}
}

func Test_Problem116(t *testing.T) {

	root := createNode(4)
	root.left = createNode(7)
	root.right = createNode(2)
	root.left.left = createNode(9)
	root.left.right = createNode(6)
	root.right.left = createNode(3)
	root.right.right = createNode(1)

	connect(root)
	fmt.Println(root.value)
	fmt.Println(root.left.next.value)
}
