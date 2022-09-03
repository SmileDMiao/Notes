package leetcode

import (
	"LeetCode/structures"
	"sort"
)

type TreeNode = structures.TreeNode

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	a1 := make([]int, 0)
	a2 := make([]int, 0)

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
