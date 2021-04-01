// 15. three sum
// question
// 在数组中找到三个数的和为0的子数组
// example
// Input: [-1, 0, 1, 2, -1, -4] Output: [-1, 0 ,1] [-1, -1, 2]

// 思路(排序+双指针)
// 先排序数组，遍历数组，当前元素nums[i],左指针i+1，右指针len(nums)-1，计算这三个值的和并与0比较大小,根据比较结果来移动指针，还要考虑到数组去重到问题

package main

import (
	"fmt"
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

func main() {
	a := []int{-1, 0, 1, 2, -1, -4}
	s := threeSum(a)
	fmt.Println(s)
}
