package leetcode

import (
	"fmt"
	"strings"
	"unicode"
)

func calculate(s string) int {
	var result int
	stack := []int{}

	// 第一个数字
	num := 0
	// 第一个肯定是数字，默认为+
	sign := '+'
	// 去掉空格
	s = strings.Replace(s, " ", "", -1)

	for i, ch := range s {
		if unicode.IsDigit(ch) {
			num = num*10 + int(ch-'0')

			// 如果出现连续数字，先找出连续的数字 num
			if i != len(s)-1 {
				continue
			}
		}

		switch sign {
		case '+':
			stack = append(stack, num)
		case '-':
			stack = append(stack, -num)
		case '*':
			newNum := stack[len(stack)-1] * num
			stack[len(stack)-1] = newNum
		case '/':
			newNum := stack[len(stack)-1] / num
			stack[len(stack)-1] = newNum
		}

		// 重置num sign
		num = 0
		sign = ch
	}

	for _, el := range stack {
		result += el
	}

	fmt.Println(result)
	return result
}
