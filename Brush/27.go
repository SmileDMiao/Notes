// 20. Remove Element
// question
// 删除数组指定元素(不要为另一个数组分配额外的空间，必须通过使用O（1）个额外内存修改输入数组来实现。)
// example
// Input nums = [3,2,2,3], val = 3, Output: 2, nums = [2,2]

// 思路(快慢指针)
// fast指针: 遍历数组元素的位置
// slow: 遍历元素不等于val的位置
// 遍历到元素不等于val时: slow++
// slow != fast时候, 左右元素交换
// 最后slow的位置就是数组元素不为val的数量,且最后等于val的元素都在数组最后

package main

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

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	removeElement(nums, 2)
}
