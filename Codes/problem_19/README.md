### 19. Remove Nth Node From End of List
> question

删除链表倒数第N个节点

> example

Input: [2, 7, 11, 15], target = 3; Output [2,7,15]


> 思路(一次循环)
设 first second都是原链表, 遍历first, 当遍历到节点数 > n的时候，
开始移动second,first second两个指针的距离是n那么遍历结束，second就位于倒数第N个节点
