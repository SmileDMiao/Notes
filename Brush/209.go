package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	left, sum, l, result := 0, 0, 0, n+1
	for i := 0; i < n; i++ {
		sum += nums[i]

		for sum >= s {
			l = i - left + 1
			result = min(result, l)
			sum -= nums[left]
			left = left + 1
		}
	}
	if result == n+1 {
		return 0
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}

	result := minSubArrayLen(s, nums)
	fmt.Println(result)
}
