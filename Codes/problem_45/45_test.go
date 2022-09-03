package leetcode

import (
	"testing"
)

func Test_Problem45(t *testing.T) {
	nums1 := []int{1, 2, 3}       //2
	nums2 := []int{2, 3, 1, 1, 4} //2
	nums3 := []int{2, 1}          //1
	nums4 := []int{3, 2, 1}       //1
	nums5 := []int{2, 3, 1}       //1
	nums6 := []int{1, 2, 1, 1, 1} //3
	nums7 := []int{2, 3, 1, 1, 4}

	jump(nums1)
	jump(nums2)
	jump(nums3)
	jump(nums4)
	jump(nums5)
	jump(nums6)
	jump(nums7)
}
