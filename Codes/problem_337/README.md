### 337. House Robber III
> question

现在房屋排列变成了二叉树, 相邻节点不能偷

> 思路(动态规划)

1. 每个节点有偷与不偷的选项
2. 偷: root.val + dp(root.right.left) + dp(root.right.right) + dp(root.left.right) + dp(root.left.left)
3. 不偷: dp(root.right) + dp(root.left)
