package leetcode

import (
	"LeetCode/structures"
	"testing"
)

func Test_Problem19(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	node := structures.Ints2List(nums)

	removeNthFromEnd(node, 2)

}
