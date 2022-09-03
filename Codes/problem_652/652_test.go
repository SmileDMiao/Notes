package leetcode

import (
	"LeetCode/structures"
	"fmt"
	"testing"
)

func Test_Problem652(t *testing.T) {
	nums := []int{1, 2, 3, 4, structures.NULL, 2, 4, structures.NULL, structures.NULL, 4}

	root := structures.Ints2TreeNode(nums)

	fmt.Println(root)

	findDuplicateSubtrees(root)
}
