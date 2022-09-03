### 503. Next Greater Element II
> question

给定一个循环数组(the next element of nums[nums.length - 1] is nums[0]), 输出每个元素的下一个更大元素, 如果不存在则输出 -1

> example

Input: nums = [1,2,1]; Output: [2,-1,2]

> 思路

遍历数组，先在元素之后的剩余数组寻找目标，找不到则在元素前面的数组寻找目标，都找不到则返回 -1
