package leetcode

import (
	"LeetCode/structures"
	"testing"
)

func Test_Problem226(t *testing.T) {
	nums := []int{4, 7, 2, 9, 6, 3, 1}
	root := structures.Ints2TreeNode(nums)
	reverse(root)

}
