// 11. Container With Most Water
// question
// 数组下标和值分别代表X Y轴, 找到最大面积
// Input: [4,3,2,1,4] Output: 16

// 思路
// 双指针法: 长度为两个数字之间距离，高度为连个数字之间的最小值
package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0

	for left < right {
		length := getSmaller(height[left], height[right])
		area := length * (right - left)
		if area > result {
			result = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return result
}

func getSmaller(first, second int) int {
	if first > second {
		return second
	}
	return first
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	maxArea(height)
}
