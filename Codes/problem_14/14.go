package leetcode

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	} else if length == 1 {
		return strs[0]
	}

	first := strs[0]
	result := ""
	for i := 1; i <= (len(first)); i++ {
		jud := true
		for _, j := range strs {
			if !strings.HasPrefix(j, first[0:i]) {
				jud = false
			}
		}
		if jud {
			result += first[(i - 1):i]
		}
	}

	return result
}
