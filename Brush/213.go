// 213. House Robber II
// question
// 打家劫舍: 不能头相邻的两家, 版本1是一个数组，这里数组变成了环形数组
// example
// Input: nums = [2, 3, 2], 7; Output: 3

// 思路(动态规划)
// 情况1. 偷了第一家就不能偷最后一家: 0 - (n-2)
// 情况2. 不偷第一家 1 - (n-1)
// 只有两件房子时候取max(m1, m2), 再多一间房子，两种情况: 偷 f(n-2) + m, 不偷 f(n-1)，取最大值
// f(n) = m1 n == 1
// f(n) = max(m1, m2) n == 2
// f(n) = max(f(n - 1), f(n - 2) + Mn)

package main

func rob(nums []int) int {
	length := len(nums)
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if length == 2 {
		return max(nums[0], nums[1])
	}

	return max(helper(nums[0:(length-1)]), helper(nums[1:length]))
}

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

func helper(nums []int) int {
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	// n-1
	dp1 := nums[0]
	// n-2
	dp2 := max(nums[1], nums[0])
	var result int
	for i := 2; i < len(nums); i++ {
		result = max(dp2, dp1+nums[i])
		dp1 = dp2
		dp2 = result
	}

	return result
}
