package leetcode

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
