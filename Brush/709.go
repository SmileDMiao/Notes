// 709. To Lower Case
// question
// 把字符串中的大写变成小写

package main

func toLowerCase(str string) string {
	downStr := []rune(str)
	sym := rune(32)

	for i, char := range str {
		if char >= 'A' && char <= 'Z' {
			downStr[i] += sym
		}
	}
	return string(downStr)
}
