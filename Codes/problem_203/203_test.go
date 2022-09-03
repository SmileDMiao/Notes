package leetcode

import (
	"LeetCode/structures"
	"testing"
)

func Test_Problem203(t *testing.T) {

	nums := []int{1, 2, 6, 3, 4, 5, 6}

	node := structures.Ints2List(nums)
	removeElements(node, 6)

}
