package leetcode

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
