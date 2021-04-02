// 20. Valid Parentheses
// question
// 字符串是否合法: 只包含()[]{}
// example
// Input "()", Output: true

// 思路
// stack存储字符, 遍历到‘{ [ (’无脑存入stack
// 遍历到‘} ] )’
// 如果stack长度为0必不合法
// 如果当前对应到正向符号不在stack最后一个(保证开关顺序)必不合法
// 合法则去掉stack最后一个
// 最后看stack长度是否为0

func isValid(s string) bool {
	parentMap := map[rune]rune{
		']': '[',
		')': '(',
		'}': '{',
	}
	var stack []rune

	for _, i := range s {
		switch i {
		case '{', '[', '(':
			stack = append(stack, i)
		case '}', ']', ')':
			if len(stack) == 0 || stack[len(stack)-1] != parentMap[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
