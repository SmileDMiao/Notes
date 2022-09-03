### 105. Construct Binary Tree from Preorder and Inorder Traversal
> question

通过前序和中序遍历构造二叉树

> example

Input preorder = [3,9,20,15,7], inorder = [9,3,15,20,7], Output: [3,9,20,null,null,15,7]

> 思路

递归构造二叉树, 前序遍历的第一个肯定是根节点，在中序遍历中找到根节点，中序遍历根节点左边是左子树，右边是右子树，然后递归构造
