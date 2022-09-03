### 1305. All Elements in Two Binary Search Trees
> question

给两个二叉搜索树,返回一个列表包含这两个二叉树的所有元素，排序从小到大

> example

Input [2,1,4], [1,0,3], Output: [0,1,1,2,3,4]

> 思路

分别中序遍历(因为是搜索二叉树)两个二叉树，然后合并排序这两个数组
