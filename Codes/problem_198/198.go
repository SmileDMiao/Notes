package leetcode

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
