### 652. Find Duplicate Subtrees
> question

给定一个二叉树，返回所有重复的子树

> example

Input: root = [1,2,3,4,null,2,4,null,null,4]; Output: [[2,4],[4]]

> 思路
问题核心是如何简单的表示一个子树且方便的比较是否重复
先续遍历二叉树，left+right+val组成字符串以表示二叉树，map以这个字符串为key，value为出现次数
