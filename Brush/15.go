// 15. three sum
// question
// 在数组中找到三个数的和为0的子数组
// example
// Input: [-1, 0, 1, 2, -1, -4] Output: [-1, 0 ,1] [-1, -1, 2]

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	var result [][]int

	if length > 0 && (nums[0] > 0 || nums[length-1] < 0) {
		return result
	}

	for i := 0; i < length-2; i++ {
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

				for left < right && nums[left] == nums[left-1] {
					left++
				}
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
