// 70. Climbing Stairs
// question
// 爬楼梯 可以爬一层或两层 爬N层楼梯有多少中爬法
// example
// Input: 2; Output: 2

// 思路(动态规划)
// 1. 第一层: 1种 记为f(1)=1 (边界)
// 2. 第二层: 2种 走2步或走两个1步 记为f(2)=2
// 3. 第三层: 3种 在第一层走2步或在第二层走1步 记为f(3)=f(1)+f(2)
// 4. 第四层: 出发点要么在第三层要么在第二层(那么到达第二层的爬法加上到达第三层的爬法就是目标值) f(n) = f(n-1) + f(n-2)
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
