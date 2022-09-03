package leetcode

func balancedStringSplit(s string) int {
	var left []string
	var right []string
	count := 0

	for _, char := range s {
		if string(char) == "L" {
			left = append(left, string(char))
		} else {
			right = append(right, string(char))
		}
		if len(right) == len(left) {
			count++
			right = right[:0]
			left = left[:0]
		}
	}
	return count
}
