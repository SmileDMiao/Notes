package leetcode

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
}
