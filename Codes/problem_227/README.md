### 227. Basic Calculator II
> question

计算字符串表达式(+-*/)

> example

Input: "3+2*2"; Output 7

> 思路(stack)

1. 去掉字符串中空格
2. 如果碰到数字， 则把数字入栈
3. 如果碰到 + - * / 找到下个数字num
4. +: num入栈
5. -: -num入栈
6. *: stack.pop * num入栈
7. /: stack.pop / num入栈
8. 最后栈内结果相加
