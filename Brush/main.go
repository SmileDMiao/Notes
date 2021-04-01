package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchRoute(root *TreeNode, target *TreeNode, cache []int) []int {
	if root == nil {
		return nil
	}
	if root.Val == target.Val {
		return append(cache, root.Val)
	}
	cache = append(cache, root.Val)
	left := make([]int, len(cache))
	right := make([]int, len(cache))
	copy(left, cache)
	copy(right, cache)
	res1 := searchRoute(root.Left, target, left)
	res2 := searchRoute(root.Right, target, right)
	if res1 != nil {
		return res1
	}
	return res2
}

func searchChild(root *TreeNode, K, deep int, res *[]int) {
	if root == nil {
		return
	}
	if K == deep {
		*res = append(*res, root.Val)
		return
	}
	searchChild(root.Left, K, deep+1, res)
	searchChild(root.Right, K, deep+1, res)
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	route := searchRoute(root, target, []int{})
	if route == nil || len(route) == 0 {
		return nil
	}
	var res []int

	for i := 0; i < len(route)-1; i++ {
		d := K - (len(route) - i - 1)
		rootNum := root.Val
		var searchNode *TreeNode
		switch {
		case root.Left == nil:
			searchNode = root.Left
			root = root.Right
		case root.Right == nil:
			searchNode = root.Right
			root = root.Left
		default:
			if route[i+1] == root.Left.Val {
				searchNode = root.Right
				root = root.Left
			} else {
				searchNode = root.Left
				root = root.Right
			}
		}

		switch {
		case d == 0:
			res = append(res, rootNum)
		case d > 0:
			searchChild(searchNode, d, 1, &res)
		default:

		}

	}
	searchChild(target, K, 0, &res)
	return res
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
