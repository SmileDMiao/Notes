// 198. House Robber
// question
// 一个数组表示有一排人家，你是一个小偷，不能偷相连的两家，问如何偷到最大金额
// example
// Input [1,2,3,1], Output: 4

// 思路(动态规划)
// 只有两件房子时候取max(m1, m2), 再多一间房子，两种情况: 偷 f(n-2) + m, 不偷 f(n-1)，取最大值
// f(n) = m1 n == 1
// f(n) = max(m1, m2) n == 2
// f(n) = max(f(n - 1), f(n - 2) + M)

package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// n - 2
	dp1 := nums[0]
	// n - 1
	dp2 := max(nums[0], nums[1])
	var result int
	for i := 2; i < len(nums); i++ {
		result = max(dp2, dp1+nums[i])
		dp1 = dp2
		dp2 = result
	}
	return result
}

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}
