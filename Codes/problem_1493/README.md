### 1493. Longest Subarray of 1's After Deleting One Element
>question

给定一个数组(数组只会包含 0 和 1)，删除一个元素，让其有一个包含1的最长子数组

> example

Input [1,1,0,1], Output: 3

> 思路

滑动窗口，维持一个区间使得区间内0的个数始终为1，求这个区间的最大长度
遍历，遇到0计数，再次循: 如果0数量大于1，nums[index]也为0，cont--,保证区间内只有1一个0
