package leetcode

func buddyStrings(A string, B string) bool {
	// base case
	if len(A) != len(B) || len(A) == 0 || len(B) == 0 {
		return false
	}

	// A B相等的情况下,字符串中必须有相等字符, map存储，如果有那么 len(s) < len(A)
	if A == B {
		s := make(map[rune]int, len(A))
		for i, char := range A {
			s[char] = i
		}

		return len(s) < len(A)
	}

	// 按照index遍历找到两个字符串不同的部分
	shortA := ""
	shortB := ""
	for i := 0; i < len(A); i++ {
		if A[i] == B[i] {
			continue
		} else {
			shortA += string(A[i])
			shortB += string(B[i])
		}
	}

	if len(shortA) == 0 {
		// 没有不同true
		return true
	} else if len(shortA) != 2 {
		// 如果不同部分大于两个字符必不符合条件
		return false
	} else if (shortA[0] == shortB[1]) && (shortA[1] == shortB[0]) {
		// 不同的俩字符调换相等返回true
		return true
	}
	return false
}
