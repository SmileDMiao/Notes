// 1493. Longest Subarray of 1's After Deleting One Element
// question
// 给定一个数组，删除一个元素，让其有一个包含1的最长子数组只(数组只会包含 0 和 1)
// example
// Input [1,1,0,1], Output: 3

// 思路
// TODO

package main

import "fmt"

func longestSubarray1(nums []int) int {
	result := -1
	count := 0
	index := 0
	for i := 0; i < len(nums); i++ {
		j := i + 1
		tmp := j
		if nums[i] == 0 {
			count++
		}
		for count == 1 && j < len(nums) {
			if nums[j] == 0 || j == len(nums)-1 {
				count--
				result = max(result, j-index)
				index = tmp
				break
			}
			j++
		}
	}

	if result == -1 {
		fmt.Println(len(nums) - 1)
	} else {
		fmt.Println(result)
	}

	return 0
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
