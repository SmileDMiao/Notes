package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	sort.Ints(nums)

	var result int
	tmp := math.MaxInt64

	for i := 0; i < length-2; i++ {
		left := i + 1
		right := length - 1

		for left < right {
			sum := nums[left] + nums[i] + nums[right]
			if sum > target {
				right--
			} else {
				left++
			}
			if tmp > abs(sum, target) {
				tmp = abs(sum, target)
				result = sum
			}
		}
	}
	return result
}

func abs(sum, target int) int {
	if sum > target {
		return sum - target
	}
	return target - sum
}
