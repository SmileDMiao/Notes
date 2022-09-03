package leetcode

type node struct {
	value int
	left  *node
	right *node
	next  *node
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
