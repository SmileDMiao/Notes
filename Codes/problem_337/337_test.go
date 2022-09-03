package leetcode

import (
	"LeetCode/structures"
	"fmt"
	"testing"
)

func Test_Problem337(t *testing.T) {
	nums := []int{3, 2, 3, 3, 1}
	root := structures.Ints2TreeNode(nums)
	fmt.Println(rob(root))
}
