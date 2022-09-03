### 1302. Deepest Leaves Sum
> question

找到二叉树最深节点的和

> example

Input [1,2,3,4,5,null,6,7,null,null,null,null,8], Output: 15

> 思路

前序遍历二叉树，遍历传入树的层级，设最大层级max，当传入层级大于max时候，result = root.Val，第一次肯定是大于的，虽然当前节点不是最深节点，但随着遍历，只有第一个最深节点变成result，然后就是当 level = max时候，result += root.Val
