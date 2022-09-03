### 224. Basic Calculator
> question

给一个表示表达式的 string，计算结果. string包含 + - ( )

> example

Input: "2-1 + 2" Output: 3

> 思路

跳过空格
遇到 "+ -" 计算结果, 重置数字和符号(标识+-)
遇到 ( 将前面的结果存入栈
遇到 ) 结果栈和sign栈pop，然后和暂时结果根据sign.pop来计算结果
最后结果+ sign*num (没有括号部分)得到最终结果
