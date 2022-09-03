### 198. House Robber
> question

一个数组表示有一排人家，你是一个小偷，不能偷相连的两家，问如何偷到最大金额

> example

Input [1,2,3,1], Output: 4

> 思路(动态规划)
设f(n)为n户人家能偷到的最大金额, M对应第n户人家的金额
只有两件房子时候取max(m1, m2), 再多一间房子，两种情况: 偷 f(n-2) + m, 不偷 f(n-1)，取最大值
f(n) = m1 n == 1
f(n) = max(m1, m2) n == 2
f(n) = max(f(n - 1), f(n - 2) + M)
