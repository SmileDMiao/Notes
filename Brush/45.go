// 45. Jump Game II
// question
// 数组的每个元素表示当前位置可以跳的最大距离，目标是到达最后的位置，求最少的跳跃次数
// example
// Input: nums = [2,3,1,1,4], target = 2, Output: 2(1 step to 1index, 3 step to last)

// 思路
// 每一次都走当前可选范围的最大值，这样可能最快的走出去，所以我们要记录当前所走的步数能到达的最远距离，
// 并且在可选范围中找到比这个值更大的最远距离，并且每次排查可选范围后，在走到最远距离时记录我们的步数。
// 最后一个元素不参与循环: [0, nums.length-1)

package main

import "fmt"

func jump(nums []int) int {
	// 记录可达最远距离
	max := 0
	// 记录步数
	step := 0
	// 寻找范围内最远距离的过程中最远距离可能会更新，所以用一个 end 变量来记录。
	end := 0

	for i := 0; i < len(nums)-1; i++ {
		max = getMax(max, nums[i]+i)

		if i == end {
			step++
			end = max
		}
		fmt.Println(end)
	}

	fmt.Println(step)
	return step
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{1, 2, 3}       //2
	nums2 := []int{2, 3, 1, 1, 4} //2
	nums3 := []int{2, 1}          //1
	nums4 := []int{3, 2, 1}       //1
	nums5 := []int{2, 3, 1}       //1
	nums6 := []int{1, 2, 1, 1, 1} //3
	nums7 := []int{2, 3, 1, 1, 4}

	jump(nums1)
	jump(nums2)
	jump(nums3)
	jump(nums4)
	jump(nums5)
	jump(nums6)
	jump(nums7)

}
