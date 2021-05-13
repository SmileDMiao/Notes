// 654. Maximum Binary Tree
// question
// 构造最大二叉树(根节点最大,然后递归构造左右节点(最大节点左右子数组))
// example
// Input: nums = [3,2,1,6,0,5], Output: [6,3,5,null,2,0,null,null,1]

// TODO
package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

func createNode(value int) *node {
	return &node{value, nil, nil}
}

func constructMaximumBinaryTree(nums []int) *node {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	var max, index int

	for i, v := range nums {
		if v >= max {
			max = v
			index = i
		}
	}

	root := createNode(max)

	root.left = constructMaximumBinaryTree(nums[0:index])
	root.right = constructMaximumBinaryTree(nums[(index + 1):])

	return root
}

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}

	root := constructMaximumBinaryTree(nums)

	fmt.Println(root.value)
	fmt.Println(root.right.value)
	fmt.Println(root.left.value)
}
