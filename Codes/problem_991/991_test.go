package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem991(t *testing.T) {
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
