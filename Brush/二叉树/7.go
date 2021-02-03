// 106通过后序和中序遍历结果构造二叉树
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

func buildTree(inorder, postorder []int) *node {
	return build(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func build(inorder []int, inStart int, inEnd int, postorder []int, postStart int, postEnd int) *node {
	if inStart > inEnd {
		return nil
	}

	rootVal := postorder[postEnd]
	var index int

	for i, k := range inorder {
		if k == rootVal {
			index = i
			break
		}
	}

	root := createNode(rootVal)

	leftSize := index - inStart

	root.left = build(inorder, inStart, index-1, postorder, postStart, postStart+leftSize-1)

	root.right = build(inorder, index+1, inEnd, postorder, postStart+leftSize, postEnd-1)

	return root
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	preorder := []int{9, 15, 7, 20, 3}

	root := buildTree(inorder, preorder)

	fmt.Println(root.value)
	fmt.Println(root.left.value)
	fmt.Println(root.right.value)
}
