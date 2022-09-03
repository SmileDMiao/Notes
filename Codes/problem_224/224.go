package leetcode

import "fmt"

func calculate(s string) int {
	cur := 0
	sign := 1
	num := 0
	res_stack := []int{}
	sign_stack := []int{}

	for i := 0; i < len(s); i++ {
		// 跳过空格
		if s[i] == ' ' {
			continue
		} else if s[i] >= '0' && s[i] <= '9' {
			// 记录数字(考虑到连续出现数字的情况)
			num = num*10 + int(s[i]-'0')
		} else if s[i] == '+' {
			// sign(标识)(+号为1 -号为-1)
			// 重置num
			// cur: 当前结果
			cur += num * sign
			num = 0
			sign = 1
		} else if s[i] == '-' {
			// cur: 暂时结果
			// 重置num
			// sign标识更改
			cur += num * sign
			num = 0
			sign = -1
		} else if s[i] == '(' {
			// (: 入栈
			// res:将前面计算的结果入栈
			// sign: 将sign入栈
			res_stack = append(res_stack, cur)
			sign_stack = append(sign_stack, sign)
			// cur重置(上面入栈了，所以要重置)
			// (内开始sign默认为1
			cur = 0
			sign = 1
		} else if s[i] == ')' {
			// 计算括号内的结果
			cur += sign * num
			// 重置num
			num = 0
			// sign_stack pop
			sign = sign_stack[len(sign_stack)-1]
			sign_stack = sign_stack[:len(sign_stack)-1]

			// res_stack pop
			// 将括号内结果和括号外结果结合计算
			cur = sign*cur + res_stack[len(res_stack)-1]
			res_stack = res_stack[:len(res_stack)-1]
		} else {
			fmt.Println("fuck")
		}
	}

	// cur: 前面计算的结果 + 最后没有括号部分的数字 = 最终结果
	result := cur + num*sign
	return result
}
