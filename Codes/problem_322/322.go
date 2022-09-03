package leetcode

func coinChange1(coins []int, amount int) int {
	// 初始化数组dp
	dp := make([]int, amount+1)
	// dp[0]=0
	dp[0] = 0

	// dp[1-amount]初始化为-1
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}
	// 循环计算dp[amount]
	for j := 1; j <= amount; j++ {
		for i := 0; i < len(coins); i++ {
			// j金额大于coins[i] 且 j - coins[i] 有最优解
			if j >= coins[i] && dp[j-coins[i]] != -1 {
				// 如果当前金额的最优解还未计算或者计算的最优解大于(dp[j-coins[i]]+1)则更新最优解
				if dp[j] == -1 || dp[j] > (dp[j-coins[i]]+1) {
					dp[j] = dp[j-coins[i]] + 1
				}
			}
		}
	}

	return dp[amount]
}

func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	for j := 1; j <= amount; j++ {
		for i := 0; i < len(coins); i++ {
			if j >= coins[i] {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
