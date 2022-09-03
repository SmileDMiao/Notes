package leetcode

import (
	"LeetCode/structures"
)

type TreeNode = structures.TreeNode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}

	if len(nums) == 0 {
		return nil
	}

	var max, index int

	for i, v := range nums {
		if v >= max {
			max = v
			index = i
		}
	}

	root := &TreeNode{Val: max, Left: nil, Right: nil}

	root.Left = constructMaximumBinaryTree(nums[0:index])
	root.Right = constructMaximumBinaryTree(nums[(index + 1):])

	return root
}
