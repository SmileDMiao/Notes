// 226. Invert Binary Tree
// question
// 翻转二叉树
// example
// Input: [4,2,7,1,3,6,9]; Output: [4,7,2,9,6,3,1]

// 思路: 每个节点要做的是 左右交换

package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

func createNode(value int) *node {
	return &node{value, nil, nil}
}

// 翻转
func reverse(root *node) *node {
	if root == nil {
		return nil
	}

	// 前序
	tmp := root.left
	root.left = root.right
	root.right = tmp

	reverse(root.right)
	reverse(root.left)

	// 后序
	// tmp := root.left
	// root.left = root.right
	// root.right = tmp

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

	reverse(root)
	fmt.Println(root.value)
	fmt.Println(root.left.value)
	fmt.Println(root.right.value)
}
