// 30. Substring with Concatenation of All Words
// question
// 给一个字符串和元素长度相同的string数组,返回子串的所有起始索引，这些子串是单词中每个单词的一次串联，顺序不限，并且没有任何中间字符
// example
// Input: s = "barfoothefoobarman", words = ["foo","bar"], Output: [0,9]

// 思路
// 遍历string, 截取数组所有元素长度的string，按照数组单个元素的长度分割string，比较截取的数组和words是否相等

package main

import (
	"reflect"
	"sort"
	"strings"
)

func findSubstring(s string, words []string) []int {
	// 所有元素组成string
	word := strings.Join(words[:], "")

	size := len(word)
	var result []int
	if len(s) == 0 || len(words) == 0 {
		return result
	}

	for i := 0; i <= (len(s) - size); i++ {
		str := s[i:(i + size)]
		if len(str) != size {
			continue
		}
		wLen := len(words[0])
		var s []string
		for i := 0; i < len(str); i = i + wLen {
			s = append(s, str[i:(i+wLen)])
		}

		sort.Strings(s)
		sort.Strings(words)
		if reflect.DeepEqual(s, words) {
			result = append(result, i)
		} else {
			continue
		}
	}
	return result
}

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}

	findSubstring(s, words)
}
