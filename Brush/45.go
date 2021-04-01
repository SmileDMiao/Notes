// todo
package main

import "fmt"

func jump(nums []int) int {
	count := 0

	tmp := nums[0]
	index := 0

	for index < (len(nums) - 1) {
		_, max := findIndex(nums, index, tmp)
		index += max
		tmp = max
		count++
		fmt.Println(max)

	}

	fmt.Println(count)
	return count
}

func findIndex(nums []int, start, end int) (int, int) {
	max := 0
	var index int
	for i, v := range nums {
		if v >= max && (start <= i) && (i < end+1) {
			max = v
			index = i
		}
	}
	return index, max
}

func main() {
	nums := []int{2, 3, 1, 1, 4}

	jump(nums)
}
