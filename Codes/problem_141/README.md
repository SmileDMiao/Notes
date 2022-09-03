### 141. Linked List Cycle
> question

给一个链表头节点，判断链表是否有环

> example

Input 3-2-0-4(指向2), Output: true

> 思路

思路1: hash
遍历链表，将节点存入hash，判断节点的下一个节点是否在hash中，在则说明有环

思路2: 双指针
快指针走两步，慢指针走一步，如果链表有环，两个指针必定相遇
