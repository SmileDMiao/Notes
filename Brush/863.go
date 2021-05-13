// 863. All Nodes Distance K in Binary Tree
// question
// 给一个二叉树，还有二叉树上的一个节点，还有一个距离K，求二叉树上的节点到目标节点距离为k的列表
// example
// Input: [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2, Output: [7,4,1]

// 思路
// TODO

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	result := make([]int, 0)

	calculate(root, root, target, K, &result)

	return result
}

func calculate(root *TreeNode, current *TreeNode, target *TreeNode, k int, result *[]int) {
	if current == nil {
		return
	}
	com := findCommonTarget(root, current, target)

	left := make([]int, 0)
	right := make([]int, 0)

	dis(com, current, &left)
	dis(com, target, &right)

	if len(left)+len(right) == k {
		*result = append(*result, current.Val)
	}

	calculate(root, current.Left, target, k, result)
	calculate(root, current.Right, target, k, result)
}

func dis(root *TreeNode, p *TreeNode, di *[]int) bool {
	if root == nil {
		return false
	}
	if root.Val == p.Val {
		return true
	}

	var found = false

	if root.Left != nil {
		found = dis(root.Left, p, di)
	}
	if !found && root.Right != nil {
		found = dis(root.Right, p, di)
	}

	if found {
		*di = append(*di, root.Val)
	}

	return found
}

func findCommonTarget(root *TreeNode, current, target *TreeNode) *TreeNode {
	if (root == nil) || root == current || root == target {
		return root
	}
	left := findCommonTarget(root.Left, current, target)
	right := findCommonTarget(root.Right, current, target)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	} else {
		return right
	}
}

func createTree(val int) *TreeNode {
	return &TreeNode{Val: val}
}
func main() {
	root := createTree(3)
	root.Left = createTree(5)
	root.Left.Left = createTree(6)
	root.Left.Right = createTree(2)
	root.Left.Right.Left = createTree(7)
	root.Left.Right.Right = createTree(4)
	root.Right = createTree(1)
	root.Right.Left = createTree(0)
	root.Right.Right = createTree(8)

	distanceK(root, root.Left, 2)
}
