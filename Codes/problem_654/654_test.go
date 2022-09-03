package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem654(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}

	root := constructMaximumBinaryTree(nums)

	fmt.Println(root.Val)
	fmt.Println(root.Right.Val)
	fmt.Println(root.Left.Val)
}
