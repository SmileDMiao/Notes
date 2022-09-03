### 114. Flatten Binary Tree to Linked List
> question

将左子树作为右子树->将原先的右子树接到当前右子树的末端

> example

Input [1,2,5,3,4,null,6], Output: [1,null,2,null,3,null,4,null,5,null,6]

> 思路
后序遍历: 每个节点做的事: 左树置空，将左树拿到右树上，将原来的右树挂到当前右树到最下面
