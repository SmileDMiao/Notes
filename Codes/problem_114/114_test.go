package leetcode

import (
	"LeetCode/structures"
	"testing"
)

func Test_Problem114(t *testing.T) {

	nums := []int{1, 2, 5, 3, 4, structures.NULL, 6}

	root := structures.Ints2TreeNode(nums)
	flatten(root)
}
