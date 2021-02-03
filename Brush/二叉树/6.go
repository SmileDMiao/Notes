// 105从前序与中序遍历构造二叉树
package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

func createNode(value int) *node {
	return &node{value, nil, nil}
}
func buildTree(preorder, inorder []int) *node {
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(preorder []int, preStart int, preEnd int, inorder []int, inStart int, inEnd int) *node {
	if preStart > preEnd {
		return nil
	}

	rootVal := preorder[preStart]
	var index int

	for i, k := range inorder {
		if k == rootVal {
			index = i
			break
		}
	}

	root := createNode(rootVal)

	leftSize := index - inStart

	root.left = build(preorder, preStart+1, preStart+leftSize, inorder, inStart, index-1)

	root.right = build(preorder, preStart+leftSize+1, preEnd, inorder, index+1, inEnd)

	return root
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)

	fmt.Println(root.value)
	fmt.Println(root.left.value)
	fmt.Println(root.right.value)
}
