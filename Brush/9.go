// 9. Palindrome Number
// question
// 正过来反过来都一样的数字返回true否则返回false
// example
// Input: 121 Output: true

// 思路
// 转成string对比反转前后的结果

package main

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
