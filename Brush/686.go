// 686. Repeated String Match
// question
// 给两个字符串 A B，重复A直到B是A的字串, 求A重复的最小次数
// example
// Input: a = "abcd", b = "cdabcdab", Output: 3

package main

// TODO

import "strings"

func repeatedStringMatch(A string, B string) int {
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
