// 322. Coin Change
// question
// 一个数组表示不同面额的金币，给一个target amount, 返回需要的最少可以拼出target amount的硬币数，可以假设每种硬币是无限的，如果不能拼出返回-1
// example
// Input: coins = [1,2,5], amount = 11; Output: 3 (5 + 5 + 1)

// 思路
// dp[amount]表示组成amount的最优解
// 特殊情况: amount = 0,dp(0)=0;amount=面值，dp(amount)=1
// 利用已知面额求出未知金额的最优解:
// dp[amount] = dp[j-coins[i]] + 1
package main

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

func main() {
	coins := []int{5, 7, 8}

	coinChange211111`eeee qqwweqwe(coins, 12)
}
