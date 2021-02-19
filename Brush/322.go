// 322
package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}
	for j := 1; j <= amount; j++ {
		for i := 0; i < len(coins); i++ {
			var m = amount + 1
			if j >= coins[i] {

				if m > min(dp[j-coins[i]]) {
					m = min(dp[j-coins[i]])
				}
				if m != (amount + 1) {
					dp[j] = m
				}
			}
		}
	}

	return dp[amount]
}

func min(a int) int {
	if a >= 0 {
		return a + 1
	}
	return 1000
}

func main() {
	coins := []int{5, 7, 8}

	result := coinChange(coins, 12)
	fmt.Println(result)
}
