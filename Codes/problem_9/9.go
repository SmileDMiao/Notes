package leetcode

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	positive_string := strconv.Itoa(x)
	number_rune := []rune(positive_string)
	for i, j := 0, len(number_rune)-1; i < j; i, j = i+1, j-1 {
		number_rune[i], number_rune[j] = number_rune[j], number_rune[i]
	}
	reverse_string := string(number_rune)

	if positive_string == reverse_string {
		return true
	} else {
		return false
	}
}
