package leetcode

import (
	"LeetCode/structures"
)

type TreeNode = structures.TreeNode

var result int
var cont int

func deepestLeavesSum(root *TreeNode) int {
	result := 0
	cont := 0
	traverse(root, cont)
	return result
}

func traverse(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if level > cont {
		cont = level
		result = root.Val
	} else if level == cont {
		result += root.Val
	}

	traverse(root.Left, level+1)
	traverse(root.Right, level+1)
}

func deepestLeavesSum1(root *TreeNode) int {
	sum := 0
	level := 0
	max := -1

	var traverse func(*TreeNode, int)
	traverse = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level > max {
			max = level
			sum = root.Val
		} else if level == max {
			sum += root.Val
		}

		traverse(root.Left, level+1)
		traverse(root.Right, level+1)
	}
	traverse(root, level)
	return sum
}
