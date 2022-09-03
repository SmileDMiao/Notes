package leetcode

import (
	"LeetCode/structures"
	"fmt"
	"testing"
)

func createNode(value int) *TreeNode {
	return &structures.TreeNode{value, nil, nil}
}

func Test_Problem1(t *testing.T) {
	root := createNode(6)
	root.Left = createNode(7)
	root.Left.Left = createNode(2)
	root.Left.Right = createNode(7)
	root.Left.Left.Left = createNode(9)
	root.Left.Right.Left = createNode(1)
	root.Left.Right.Right = createNode(4)

	root.Right = createNode(8)
	root.Right.Left = createNode(1)
	root.Right.Right = createNode(3)
	root.Right.Right.Right = createNode(5)

	fmt.Println(deepestLeavesSum(root))
	fmt.Println(deepestLeavesSum1(root))
}
