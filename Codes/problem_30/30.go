package leetcode

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
