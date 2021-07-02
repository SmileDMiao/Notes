// 503. Next Greater Element II
// question
// 给定一个循环数组(the next element of nums[nums.length - 1] is nums[0]), 输出每个元素的下一个更大元素, 如果不存在则输出 -1
// example
// Input: nums = [1,2,1]; Output: [2,-1,2]

// 思路
// 遍历数组，先在元素之后的剩余数组寻找目标，找不到则在元素前面的数组寻找目标，都找不到则返回 -1

package main

func nextGreaterElements(nums []int) []int {
	// 定义结果
	result := make([]int, len(nums))

	// 循环
	for i := 0; i < len(nums); i++ {
		result[i] = findBigger(nums, i)
	}

	return result
}

func findBigger(nums []int, index int) int {
	// index: 当前数字
	// 当前数字往后找到了大于当前数字则返回那个大的数字
	for i := index; i < len(nums); i++ {
		if nums[i] > nums[index] {
			return nums[i]
		}
	}

	// 在当前数字前寻找更大的数字
	for i := 0; i < index; i++ {
		if nums[i] > nums[index] {
			return nums[i]
		}
	}

	// 找不到返回-1
	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 3}

	nextGreaterElements(nums)
}
