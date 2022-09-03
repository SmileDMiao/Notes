package leetcode

import (
	"LeetCode/structures"
	"testing"
)

func Test_Problem817(t *testing.T) {
	G := []int{0, 1, 3, 4, 7, 9, 10}
	nums := []int{0, 1, 2, 3, 4, 6, 7, 8, 9, 10, 10}
	head := structures.Ints2List(nums)

	numComponents(head, G)
}
