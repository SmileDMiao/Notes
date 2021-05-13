// TODO
func nextGreaterElements(nums []int) []int {
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		result[i] = findBigger(nums, i)
	}

	return result
}

func findBigger(nums []int, index int) int {
	for i := index; i < len(nums); i++ {
		if nums[i] > nums[index] {
			return nums[i]
		}
	}

	for i := 0; i < index; i++ {
		if nums[i] > nums[index] {
			return nums[i]
		}
	}
	return -1
}
