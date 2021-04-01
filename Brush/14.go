// 14. Longest Common Prefix
// question
// 最长公共前缀
// example
// Input: ["flower","flow","flight"] Output: "fl"

// 思路
// 公共前缀, 两次遍历，第一次遍历第一个元素，第二次遍历剩余数组元素
package main

import (
	"fmt"
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
		if jud == true {
			result += first[(i - 1):i]
		}
	}
	fmt.Println(result)
	return result
}
