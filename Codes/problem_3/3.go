package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	sliceString := []rune(s)
	length, maxLength := 0, 0
	m := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		v, ok := m[sliceString[i]]
		if ok {
			d := i - v
			if d > length {
				length++
			} else {
				length = d
			}
		} else {
			length++
		}
		if length > maxLength {
			maxLength = length
		}
		m[sliceString[i]] = i
	}
	fmt.Println(maxLength)
	return maxLength
}

func main() {
	s := "arabcacfr"
	lengthOfLongestSubstring(s)
}
