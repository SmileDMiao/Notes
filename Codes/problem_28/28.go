package leetcode

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
