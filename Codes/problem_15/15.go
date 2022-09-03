package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	var result [][]int

	// 这几种情况直接返回空
	if length > 0 && (nums[0] > 0 || nums[length-1] < 0) {
		return result
	}

	for i := 0; i < length-2; i++ {
		// 不是第一个元素且和前面的相等则跳过(去重)
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := length - 1

		for left < right {
			if nums[left]+nums[right]+nums[i] == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// 去重
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				// 去重
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
			if (left < right) && (nums[i]+nums[left]+nums[right]) > 0 {
				right--
			}
			if (left < right) && (nums[i]+nums[left]+nums[right] < 0) {
				left++
			}
		}
	}
	return result
}
