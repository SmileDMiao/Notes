// 114. Flatten Binary Tree to Linked List
// question
// 将左子树作为右子树->将原先的右子树接到当前右子树的末端
// example
// Input [1,2,5,3,4,null,6], Output: [1,null,2,null,3,null,4,null,5,null,6]

// 思路
// 后序遍历: 每个节点做的事: 左树置空，将左树拿到右树上，将原来的右树挂到当前右树到最下面

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
