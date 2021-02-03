// 70
package main

import "fmt"

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func main() {
	result := climbStairs(2)
	fmt.Println(result)
}
