package leetcode

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
