### 1221. Split a String in Balanced Strings
> question

balance string包含相同数量的 L R，分割字符串，有多少个balance string

> example

Input: "RLRRLLRLRL"; Output: 4

> 思路

遍历字符串，L放左边，R放右边，每次遍历判断左右是否长度一样，如果一样则count+1然后清空左右数组
