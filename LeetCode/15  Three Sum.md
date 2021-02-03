```go
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1

		if nums[i] > 0 {
			return result
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				left++
				right--
				result = append(result, []int{nums[i], nums[left], nums[right]})
			}
			if nums[i]+nums[left]+nums[right] > 0 {
				right--
			}
			if nums[i]+nums[left]+nums[right] < 0 {
				left++
			}
		}
	}
	return result
}
```