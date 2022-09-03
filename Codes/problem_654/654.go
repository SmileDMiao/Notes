package leetcode

import (
	"LeetCode/structures"
)

type TreeNode structures.TreeNode

func constructMaximumBinaryTree(nums []int) *structures.TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	var max, index int

	for i, v := range nums {
		if v >= max {
			max = v
			index = i
		}
	}

	root := &structures.TreeNode{max, nil, nil}

	root.Left = constructMaximumBinaryTree(nums[0:index])
	root.Right = constructMaximumBinaryTree(nums[(index + 1):])

	return root
}
