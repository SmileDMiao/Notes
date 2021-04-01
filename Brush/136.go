// 136. Single Number
// question
// 找到数组中只出现一次的数字
// example
// Input [2,2,1], Output: 1

// 思路1(map)
// map保存数组中的数字以及对应的出现次数，然后找到只出现一次的

package main

func singleNumber1(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		_, ok := m[v]
		if !ok {
			m[v]++
		} else {
			if m[v] == 1 {
				m[v]++
			}
		}
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

func singleNumber2(nums []int) int {
	result := 0
	for _, v := range nums {
		result = result ^ v
	}
	return result
}

func main() {
	nums := []int{2, 2, 3, 3, 4, 5, 5}

	singleNumber2(nums)
}
