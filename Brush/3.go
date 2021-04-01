// 3. Longest Substring Without Repeating Characters
// question
// 不重复的最长字串
// example
// Input: "abcabcbb"; Output: 3

// 思路(动态规划)
// f(i): 以第i个字符为结尾的不包含重复字符的子字符串的最长长度。
// 如果第i个字符之前没有出现过: f(i) = f(i-1) + 1
// 如果第i个字符之前出现过: 设d为当前字符与上次出现字符之间第距离
// 当 d <= f(i)时: f(i) = d
// 当 d > f(i)时: f(i) = f(i-1) + 1

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	sliceString := []rune(s)
	length, maxLength := 0, 0
	m := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		v, ok := m[sliceString[i]]
		if ok {
			d := i - v
			if d > length {
				length++
			} else {
				length = d
			}
		} else {
			length++
		}
		if length > maxLength {
			maxLength = length
		}
		m[sliceString[i]] = i
	}
	fmt.Println(maxLength)
	return maxLength
}

func main() {
	s := "arabcacfr"
	lengthOfLongestSubstring(s)
}
