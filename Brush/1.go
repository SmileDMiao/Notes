// 1. TwoSum
// question
// 给一个数组，一个数字，找到数组中两个元素的和为目标数字的下标
// example
// Input: [2, 7, 11, 15], target = 9; Output [0, 1]

// 思路
// 定义一个map, 循环数组, 第一次先将值作为key，索引作为value存入map, 后面找的时候，看targe - nums[i]是否在map中, 如果在则写入结果
package main

import "fmt"

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
