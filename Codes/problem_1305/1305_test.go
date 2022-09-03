package leetcode

import (
	"LeetCode/structures"
	"fmt"
	"testing"
)

func Test_Problem1305(t *testing.T) {
	a1 := []int{2, 1, 4}
	a2 := []int{1, 0, 3}

	tree1 := structures.Ints2TreeNode(a1)
	tree2 := structures.Ints2TreeNode(a2)

	fmt.Println(getAllElements(tree1, tree2))
}
