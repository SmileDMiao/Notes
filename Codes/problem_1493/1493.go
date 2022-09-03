package leetcode

func longestSubarray1(nums []int) int {
	// 结果
	result := 0
	// 0计数
	count := 0
	// 区间
	index := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		}

		for count > 1 && index < len(nums) {
			if nums[index] == 0 {
				count--
			}
			index++
		}

		result = max(result, i-index)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 思路
// 遍历数组，遇到1标记下, 然后往后遍历剩下的数组，遇到不是1的再次标记，遇到不是1的第二次停止，记录max
func longestSubarray2(nums []int) int {
	var max int = 0
	for i := 0; i < len(nums); i++ {
		var count int
		if nums[i] == 1 {
			count = 1
		} else {
			continue
		}
		var sym bool = false
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == 1 {
				count++
			} else {
				if sym == true {
					break
				} else {
					sym = true
				}
			}
		}
		if count > max {
			max = count
		}
	}

	if max == len(nums) {
		return len(nums) - 1
	}

	return max
}
