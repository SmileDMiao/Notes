package leetcode

import "strings"

func repeatedStringMatch(A string, B string) int {
	// base case
	if len(A) == 0 && len(B) == 0 {
		return 1
	} else if len(A) == 0 || len(B) == 0 {
		return 0
	}

	S := A
	count := 0

	limit := len(B)/len(A) + 1

	for count <= limit {
		if strings.Index(S, B) > -1 {
			return count + 1
		}
		S += A
		count++
	}

	return -1
}
