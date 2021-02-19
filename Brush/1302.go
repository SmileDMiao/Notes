// 1302. Deepest Leaves Sum
// question
// 找到二叉树最深节点的和
// example
// Input [1,2,3,4,5,null,6,7,null,null,null,null,8], Output: 15

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result int
var cont int
var maxCont int = -1

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
}
