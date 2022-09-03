package leetcode

import "fmt"

func jump(nums []int) int {
	// 记录可达最远距离
	max := 0
	// 记录步数
	step := 0
	// 寻找范围内最远距离的过程中最远距离可能会更新，所以用一个 end 变量来记录。
	end := 0

	for i := 0; i < len(nums)-1; i++ {
		max = getMax(max, nums[i]+i)

		if i == end {
			step++
			end = max
		}
		fmt.Println(end)
	}

	fmt.Println(step)
	return step
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
