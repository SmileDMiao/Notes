### 20. Valid Parentheses
> question

字符串是否合法: 只包含()[]{} 需要闭合

> example

Input "()", Output: true

> 思路

stack存储字符, 遍历到 { [ ( 无脑存入stack
遍历到 } ] )
如果stack长度为0必不合法
如果当前对应到正向符号不在stack最后一个(保证开关顺序)必不合法
合法则去掉stack最后一个
最后看stack长度是否为0
