// 16. 3sum closest
// question
// 找到三个数的和离目标和最近
// example
// Input [1,1,1,0] -100, Output: 2

// 思路
// 排序数组，tmp记录距离，result记录正确结果，还是双指针的做法

package main

import (
	"fmt"
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

func main() {
	nums := []int{1, 1, 1, 0}
	fmt.Println(threeSumClosest(nums, -100))
}
