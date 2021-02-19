package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

var result = make(map[*TreeNode]int)

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if v, ok := result[root]; ok {
		return v
	}
	do := root.Val
	if root.Right != nil {
		do += rob(root.Right.Left) + rob(root.Right.Right)
	}
	if root.Left != nil {
		do += rob(root.Left.Left) + rob(root.Left.Right)
	}

	not := rob(root.Left) + rob(root.Right)

	s := max(do, not)

	result[root] = s
	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := createNode(3)
	root.Left = createNode(2)
	root.Right = createNode(3)
	root.Left.Right = createNode(3)
	root.Right.Right = createNode(1)
	fmt.Println(root)
	fmt.Println(rob(root))
}
