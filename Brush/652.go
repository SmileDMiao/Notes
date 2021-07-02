// 652. Find Duplicate Subtrees
// question
// 给定一个二叉树，返回所有重复的子树
// example
// Input: root = [1,2,3,4,null,2,4,null,null,4]; Output: [[2,4],[4]]

// 思路
// 问题核心是如何简单的表示一个子树且方便的比较是否重复
// 先续遍历二叉树，left+right+val组成字符串以表示二叉树，map以这个字符串为key，value为出现次数

package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
