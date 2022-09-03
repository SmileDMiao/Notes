package leetcode

import (
	"testing"
)

func Test_Problem1493(t *testing.T) {
	nums1 := []int{1, 1, 0, 1}
	nums2 := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	nums3 := []int{1, 1, 1}
	nums4 := []int{1, 1, 0, 0, 1, 1, 1, 0, 1}
	nums5 := []int{0, 0, 0}

	// 3
	longestSubarray1(nums1)
	longestSubarray2(nums1)

	// 5
	longestSubarray1(nums2)
	// 2
	longestSubarray1(nums3)
	// 4
	longestSubarray1(nums4)
	// 0
	longestSubarray1(nums5)
}
