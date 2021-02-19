1. 用字符串将树表示出来
2. 将数量缓存进map

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var list []*TreeNode
var result map[string]int

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	list = make([]*TreeNode, 0)
	result = make(map[string]int)
	traverse(root)
	return list
}
func traverse(root *TreeNode) string {
	if root == nil {
		return "-"
	}
	left := traverse(root.Left)
	right := traverse(root.Right)

	subTree := left + "," + right + "," + strconv.Itoa(root.Val)

	if result[subTree] == 1 {
		list = append(list, root)
	}
	result[subTree]++

	return subTree
}
```