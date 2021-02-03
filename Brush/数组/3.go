// 494目标和

package main

func findTargetSumWay(nums []int, target int) int {
	
}

def backtrack(nums, i):
    if i == len(nums):
        if 达到 target:
            result += 1
        return

    for op in { +1, -1 }:
        选择 op * nums[i]
        # 穷举 nums[i + 1] 的选择
        backtrack(nums, i + 1)
        撤销选择
