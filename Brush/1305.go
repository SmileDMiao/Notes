// 1305. All Elements in Two Binary Search Trees
// question
// 给两个二叉搜索树,返回一个列表包含这两个二叉树的所有元素，排序从小到大
// example
// Input [2,1,4], [1,0,3], Output: [0,1,1,2,3,4]

// 思路
// 分别中序遍历(因为是搜索二叉树)两个二叉树，然后合并排序这两个数组

package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var a1 []int
var a2 []int

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	a1 = make([]int, 0)
	a2 = make([]int, 0)

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
