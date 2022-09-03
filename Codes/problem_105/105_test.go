package leetcode

import (
	"LeetCode/structures"
	"fmt"
	"testing"
)

func Test_Problem105(t *testing.T) {
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}

	v := structures.Tree2Preorder(buildTree(pre, in))
	fmt.Println(v)
}
