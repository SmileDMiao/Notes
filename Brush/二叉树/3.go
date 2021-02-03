// 116填充二叉树节点的右侧指针
// 每个节点要做的是: 左边绑定 右边绑定 left.right绑定right.left
package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
	next  *node
}

func createNode(value int) *node {
	return &node{value, nil, nil, nil}
}

func connect(node *node) *node {
	if node == nil {
		return nil
	}
	connectTwoNode(node.left, node.right)
	return node
}

func connectTwoNode(node1 *node, node2 *node) {
	if node1 == nil || node2 == nil {
		return
	}

	// 前序遍历位置
	// 将传入的两个节点连接
	node1.next = node2

	// 连接相同父节点的两个子节点
	connectTwoNode(node1.left, node1.right)
	connectTwoNode(node2.left, node2.right)
	// 连接跨越父节点的两个子节点
	connectTwoNode(node1.right, node2.left)
}

func main() {
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
