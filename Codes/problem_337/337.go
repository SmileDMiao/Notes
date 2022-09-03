package leetcode

import (
	"LeetCode/structures"
)

type TreeNode = structures.TreeNode

var result = make(map[*TreeNode]int)

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 查看当前节点有没有偷过
	if v, ok := result[root]; ok {
		return v
	}

	do := root.Val

	// 当前节点偷
	if root.Right != nil {
		do += rob(root.Right.Left) + rob(root.Right.Right)
	}
	if root.Left != nil {
		do += rob(root.Left.Left) + rob(root.Left.Right)
	}

	// 当前节点不偷
	not := rob(root.Left) + rob(root.Right)

	// 偷与不偷取大的那个
	s := max(do, not)

	// 存储节点是否有没有偷
	result[root] = s
	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
