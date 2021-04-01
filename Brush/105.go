// 105. Construct Binary Tree from Preorder and Inorder Traversal
// question
// 通过前序和中序遍历构造二叉树
// example
// Input preorder = [3,9,20,15,7], inorder = [9,3,15,20,7], Output: [3,9,20,null,null,15,7]

// 思路
// 递归构造二叉树, 前序遍历的第一个肯定是根节点，在中序遍历中找到根节点，中序遍历根节点左边是左子树，右边是右子树，然后递归构造

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func main() {
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}

	buildTree(pre, in)

}
