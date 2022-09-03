package leetcode

import "LeetCode/structures"

type TreeNode = structures.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(pre []int, prestart int, preend int, in []int, instart int, inend int) *TreeNode {
	if prestart > preend {
		return nil
	}

	rootVal := pre[prestart]

	var index int
	for i := instart; i <= inend; i++ {
		if in[i] == rootVal {
			index = i
			break
		}
	}

	leftsize := index - instart

	root := &TreeNode{
		Val: rootVal,
	}

	root.Left = build(pre, prestart+1, prestart+leftsize, in, instart, index-1)
	root.Right = build(pre, prestart+leftsize+1, preend, in, index+1, inend)

	return root
}
