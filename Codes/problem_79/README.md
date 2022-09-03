### 79. Word Search
> question

给一个二维数组(矩阵)和一个string，判断这个矩阵中是否包含这个string(矩阵中的元素连续起来出现组成string)

> example

Input: [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], "ABCCED"; Output: true


> 思路

遍历二维数组，DFS(深度优先搜索)，查看每个元素可达的相邻元素
