### 203. Remove Linked List Elements
> question

删除链表中value为指定值的节点

> example

Input [1,2,6,3,4,5,6], val = 6, Output: [1,2,3,4,5]

> 思路

添加一个空链表头，这样每个节点的操作就一致了(包含假节点, 每个节点判断下一节点的值)
