// 遍历二叉树
package main

import "fmt"

// Node struct
type node struct {
	value int
	left  *node
	right *node
}

// reateNode 创建Node
func createNode(value int) *node {
	return &node{value, nil, nil}
}

// 前序遍历
func (node *node) preOrder(n *node) {
	if n != nil {
		fmt.Printf("%d ", n.value)
		node.preOrder(n.left)
		node.preOrder(n.right)
	}
}

// 中序遍历
func (node *node) inOrder(n *node) {
	if n != nil {
		node.inOrder(n.left)
		fmt.Printf("%d ", n.value)
		node.inOrder(n.right)
	}
}

// 后续遍历
func (node *node) postOrder(n *node) {
	if n != nil {
		node.postOrder(n.left)
		node.postOrder(n.right)
		fmt.Printf("%d ", n.value)
	}
}

func main() {
	//创建一颗树
	root := createNode(5)
	root.left = createNode(2)
	root.right = createNode(4)
	root.left.right = createNode(7)
	root.left.right.left = createNode(6)
	root.right.left = createNode(8)
	root.right.right = createNode(9)

	fmt.Println("前序遍历")
	root.preOrder(root)

	fmt.Println("\n中序遍历")
	root.inOrder(root)

	fmt.Println("\n后序遍历")
	root.postOrder(root)
}
