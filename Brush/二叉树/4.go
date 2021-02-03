// 114将二叉树展开为链表
// 将左子树作为右子树->将原先的右子树接到当前右子树的末端
package main

import (
	"fmt"
)

type node struct {
	value int
	left  *node
	right *node
}

func createNode(value int) *node {
	return &node{value, nil, nil}
}

func flatten(root *node) *node {
	if root == nil {
		return nil
	}
	flatten(root.left)
	flatten(root.right)

	left := root.left
	right := root.right

	root.left = nil
	root.right = left

	for root.right != nil {
		root = root.right
	}
	root.right = right
	return root
}

func main() {
	root := createNode(4)
	root.left = createNode(7)
	root.right = createNode(2)
	root.left.left = createNode(9)
	root.left.right = createNode(6)
	root.right.left = createNode(3)
	root.right.right = createNode(1)

	flatten(root)
	fmt.Println(root.value)
	fmt.Println(root.right.value)
	fmt.Println(root.right.right.value)
	fmt.Println(root.right.right.right.value)
	fmt.Println(root.right.right.right.right.value)
	fmt.Println(root.right.right.right.right.right.value)
	fmt.Println(root.right.right.right.right.right.right.value)
}
