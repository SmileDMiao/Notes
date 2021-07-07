// 1493. Longest Subarray of 1's After Deleting One Element
// question
// 给定一个数组(数组只会包含 0 和 1)，删除一个元素，让其有一个包含1的最长子数组
// example
// Input [1,1,0,1], Output: 3

// 思路(滑动窗口，维持一个区间使得区间内0的个数始终为1，求这个区间的最大长度)
// 遍历，遇到0计数，再次循: 如果0数量大于1，nums[index]也为0，cont--,保证区间内只有1一个0

package main

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

func main() {
	nums1 := []int{1, 1, 0, 1}
	nums2 := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	nums3 := []int{1, 1, 1}
	nums4 := []int{1, 1, 0, 0, 1, 1, 1, 0, 1}
	nums5 := []int{0, 0, 0}

	// 3
	longestSubarray1(nums1)
	// 5
	longestSubarray1(nums2)
	// 2
	longestSubarray1(nums3)
	// 4
	longestSubarray1(nums4)
	// 0
	longestSubarray1(nums5)
}
