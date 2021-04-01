// 18. four sum
// question
// 找到4个数和为target
// example
// Input: [1,0,-1,0,-2,2], target = 0 Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

// 思路(排序+双指针 这里是两层循环)
// 先排序数组，两层循环数组，这里的思路是在第二层循环中寻找三个数的和为:target - nums[i]的,转化为threeSum问题

package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	var result [][]int

	for i := 0; i < len(nums)-1; i++ {
		sum := target - nums[i]

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-1; j++ {
			left := j + 1
			right := len(nums) - 1

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for left < right {
				if nums[left]+nums[right]+nums[j] == sum {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--

					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				}
				if (left < right) && (nums[j]+nums[left]+nums[right]) > sum {
					right--
				}
				if (left < right) && (nums[j]+nums[left]+nums[right] < sum) {
					left++
				}
			}
		}
	}
	return result
}

func main() {
	nums := []int{-2, -1, -1, 1, 1, 2, 2}

	fourSum(nums, 0)

	fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
}
