// 7. Reverse Integer
// question
// 反转一个数字
// example
// Input: 123; Output: 321

// 思路
// 依次获取数字的最后一个数(%10), 原数字/10

package main

func reverse(x int) int {
	const MaxInt32 = int(2147483647)
	const MinInt32 = int(-2147483648)
	var result int

	for x != 0 {
		tmp := x % 10

		result *= 10
		result += tmp
		x /= 10
	}

	if MaxInt32 < result || result < MinInt32 {
		return 0
	}

	return result
}

func main() {
	reverse(-321)
}
