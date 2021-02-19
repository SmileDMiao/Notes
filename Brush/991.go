package main

import "fmt"

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

func main() {
	a, b := 2, 3
	fmt.Println(brokenCalc(a, b))
	a = 5
	b = 8
	fmt.Println(brokenCalc(a, b))

	a = 3
	b = 10
	fmt.Println(brokenCalc(a, b))

	a = 1
	b = 1000000000
	fmt.Println(brokenCalc(a, b))
}
