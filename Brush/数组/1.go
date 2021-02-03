// 1TwoSum
package main

import "fmt"

// question: 给一个数组，一个数字，找到数组中两个元素的和为目标数字的下标
// [2, 7, 11, 15], target = 9 return [0, 1]

// answer: 定义一个map，循环数组，在map中找是否存在key: target - nums[i], 找到就跳出，找不到map[nums[i]] = i
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}

	fmt.Println(twoSum(nums, 9))
}
