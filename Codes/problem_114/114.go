package leetcode

import (
	"LeetCode/structures"
)

type TreeNode = structures.TreeNode

func flatten(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	for root.Right != nil {
		root = root.Right
	}
	root.Right = right
	return root
}
