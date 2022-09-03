package leetcode

import (
	"LeetCode/structures"
	"strconv"
)

type TreeNode = structures.TreeNode

// 存储结果
var list []*TreeNode

// key表示为二叉子树 值表示为出现次数
var result map[string]int

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	list = make([]*TreeNode, 0)
	result = make(map[string]int)

	traverse(root)
	return list
}

func traverse(root *TreeNode) string {
	if root == nil {
		return "-"
	}
	left := traverse(root.Left)
	right := traverse(root.Right)

	// 表示二叉树
	subTree := left + "," + right + "," + strconv.Itoa(root.Val)

	// 先判断后计数
	// if result[subTree] == 1 {
	// 	list = append(list, root)
	// }

	// 计数
	// result[subTree]++

	// return subTree

	// 先计数后判断
	result[subTree]++

	//
	if result[subTree] == 2 {
		list = append(list, root)
	}

	return subTree
}
