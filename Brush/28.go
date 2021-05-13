// 20. Implement strStr()
// question
// 返回字串在string中第一次出现的index
// example
// Input: haystack = "hello", needle = "ll", Output: 2

package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	n := []rune(needle)

	for i, v := range haystack {
		if (i + len(n)) > len(haystack) {
			break
		}
		if v != n[0] {
			continue
		} else if haystack[i:i+len(n)] == needle {
			return i
		}
	}
	return -1
}
