### 322. Coin Change
> question

一个数组表示不同面额的金币，给一个target amount, 返回需要的最少可以拼出target amount的硬币数，可以假设每种硬币是无限的，如果不能拼出返回-1

> example

Input: coins = [1,2,5], amount = 11; Output: 3 (5 + 5 + 1)

> 思路(动态规划)

1. dp[amount]表示组成amount的最优解
2. 特殊情况: amount = 0,dp(0)=0;amount=面值，dp(amount)=1
3. 利用已知面额求出未知金额的最优解: dp[amount] = dp[j-coins[i]] + 1
