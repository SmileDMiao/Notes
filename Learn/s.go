package main

import (
	"fmt"
	"unsafe"
)

func longestSubarray(nums []int) int {
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
		fmt.Println(max - 1)

		return len(nums) - 1
	}

	fmt.Println(max)

	return max
}
func main() {
	var a []int
	a = []int{1, 1, 0, 1}
	longestSubarray(a)

	a = []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	longestSubarray(a)

	a = []int{1, 1, 1}
	longestSubarray(a)

	a = []int{1, 1, 0, 0, 1, 1, 1, 0, 1}
	longestSubarray(a)
	a = []int{0, 0, 0}
	longestSubarray(a)

	a = []int{1, 2, 3, 4, 1, 2, 1, 1}
	longestSubarray(a)

	fmt.Println(unsafe.Sizeof(struct{}{}))


}
