package leetcode

import (
	"LeetCode/structures"
)

type TreeNode = structures.TreeNode

// 翻转
func reverse(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 前序
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	reverse(root.Right)
	reverse(root.Left)

	// 后序
	// tmp := root.left
	// root.left = root.right
	// root.right = tmp

	return root
}
