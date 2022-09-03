package leetcode

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
