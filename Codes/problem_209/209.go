package leetcode

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
