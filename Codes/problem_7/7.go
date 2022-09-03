package leetcode

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
