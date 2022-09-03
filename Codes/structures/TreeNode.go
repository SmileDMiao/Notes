package structures

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方便生成测试数据
var NULL = -1 << 63

// 通过数组生成二叉树
// 按照层级顺序(LeetCode网站给出的顺序)
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// 前序遍历二叉树
func Tree2Preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := []int{root.Val}
	res = append(res, Tree2Preorder(root.Left)...)
	res = append(res, Tree2Preorder(root.Right)...)

	return res
}

// 中序遍历二叉树
func Tree2Inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := Tree2Inorder(root.Left)
	res = append(res, root.Val)
	res = append(res, Tree2Inorder(root.Right)...)

	return res
}

// 后续遍历二叉树
func Tree2Postorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := Tree2Postorder(root.Left)
	res = append(res, Tree2Postorder(root.Right)...)
	res = append(res, root.Val)

	return res
}
