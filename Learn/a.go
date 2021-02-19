package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

var a1 []int
var a2 []int

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	a1 = make([]int, 0)
	a2 = make([]int, 0)

	traverse(root1, &a1)
	traverse(root2, &a2)

	a1 = append(a1, a2...)
	sort.Ints(a1)
	return a1
}

func traverse(root *TreeNode, a *[]int) {
	if root == nil {
		return
	}

	traverse(root.Left, a)
	*a = append(*a, root.Val)
	traverse(root.Right, a)
}
func main() {
	node1 := createNode(0)
	node1.Left = createNode(-10)
	node1.Right = createNode(10)
	node1.Left.Left = createNode(-11)
	node1.Left.Right = createNode(-8)

	node2 := createNode(5)
	node2.Left = createNode(1)
	node2.Right = createNode(7)
	node2.Left.Left = createNode(0)
	node2.Left.Right = createNode(2)

	fmt.Println(getAllElements(node1, node2))

}
