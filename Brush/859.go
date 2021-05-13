// 859. Buddy Strings
// question
// 给两个字符串A，B，如果在A中交换两个字符能得到B return true 否则 return false
// exmaple
// Input: a = "ab", b = "ba", Output: true

// TODO
package main

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) || len(A) == 0 || len(B) == 0 {
		return false
	}

	if A == B {
		s := make(map[rune]int, len(A))
		for i, char := range A {
			s[char] = i
		}
		return len(s) < len(A)
	}

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
		return true
	} else if len(shortA) != 2 {
		return false
	} else if (shortA[0] == shortB[1]) && (shortA[1] == shortB[0]) {
		return true
	}
	return false
}
