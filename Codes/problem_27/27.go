package leetcode

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			if fast != slow {
				nums[fast], nums[slow] = nums[slow], nums[fast]
			}
			slow++
		}
	}

	return slow
}
