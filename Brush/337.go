// 337. House Robber III
// question
// 现在房屋排列变成了二叉树, 相邻节点不能偷

// 思路:
// 1. 每个节点有偷与不偷的选项
// 2. 偷: root.val + dp(root.right.left) + dp(root.right.right) + dp(root.left.right) + dp(root.left.left)
// 3. 不偷: dp(root.right) + dp(root.left)

// TODO

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

var result = make(map[*TreeNode]int)

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if v, ok := result[root]; ok {
		return v
	}

	do := root.Val
	if root.Right != nil {
		do += rob(root.Right.Left) + rob(root.Right.Right)
	}
	if root.Left != nil {
		do += rob(root.Left.Left) + rob(root.Left.Right)
	}

	not := rob(root.Left) + rob(root.Right)

	s := max(do, not)

	result[root] = s
	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := createNode(3)
	root.Left = createNode(2)
	root.Right = createNode(3)
	root.Left.Right = createNode(3)
	root.Right.Right = createNode(1)
	fmt.Println(root)
	fmt.Println(rob(root))
}
