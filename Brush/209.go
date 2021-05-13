// 209. Minimum Size Subarray Sum
// question
// 一个含有n个正整数的数组和一个正整数s，找出该数组中满足其和 ≥ s的长度最小的连续子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
// example
// Input: nums = [2,3,1,2,4,3], 7; Output: 2

// 思路(滑动窗口)
// 遍历数组记录和，直到和 >= s, 此时移动左指针, sum >= s继续右移，条件不成立，继续往后遍历数组

package main

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)

	// base case
	if n == 0 {
		return 0
	}

	// left: 左边指针
	// sum: 记录区间和
	// result: 记录符合条件子数组长度
	left, sum, result := 0, 0, n+1
	for i := 0; i < n; i++ {
		sum += nums[i]

		for sum >= s {
			l := i - left + 1
			result = min(result, l)
			sum -= nums[left]
			left = left + 1
		}
	}
	// 没有符合条件的子数组
	if result == n+1 {
		return 0
	}

	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}

	minSubArrayLen(s, nums)
}
