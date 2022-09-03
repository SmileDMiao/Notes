### 116. Populating Next Right Pointers in Each Node
> question

填充二叉树节点的右侧指针

> example

Input: "abcabcbb"; Output: 3

> 思路

每个节点要做的是: 左边绑定 右边绑定 left.right绑定right.left
两种情况: 1连接相同父节点的两个子节点; 2连接跨越父节点的两个子节点
