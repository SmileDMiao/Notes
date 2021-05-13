// 1302. Deepest Leaves Sum
// question
// 找到二叉树最深节点的和
// example
// Input [1,2,3,4,5,null,6,7,null,null,null,null,8], Output: 15

// 思路
// 前序遍历二叉树，遍历传入树的层级，设最大层级max，当传入层级大于max时候，result = root.Val，第一次肯定是大于的，虽然当前节点不是最深节点，但随着
// 遍历，只有第一个最深节点变成result，然后就是当 level = max时候，result += root.Val

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result int
var cont int

func deepestLeavesSum(root *TreeNode) int {
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

func createNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

func main() {
	root := createNode(6)
	root.Left = createNode(7)
	root.Left.Left = createNode(2)
	root.Left.Right = createNode(7)
	root.Left.Left.Left = createNode(9)
	root.Left.Right.Left = createNode(1)
	root.Left.Right.Right = createNode(4)

	root.Right = createNode(8)
	root.Right.Left = createNode(1)
	root.Right.Right = createNode(3)
	root.Right.Right.Right = createNode(5)

	fmt.Println(deepestLeavesSum(root))
	fmt.Println(deepestLeavesSum1(root))
}
