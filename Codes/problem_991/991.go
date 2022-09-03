package leetcode

func brokenCalc(X int, Y int) int {
	if X > Y {
		return X - Y
	}

	if X == Y {
		return 0
	}

	count := 0

	for Y > X {
		count++
		if Y%2 != 0 {
			Y++
			continue
		}
		Y /= 2
	}

	return X - Y + count
}
