// 991. Broken Calculator
// question
// 一个坏的计算器, 只能进行两种操作(-1 和 *2)，给两个数字 x,y。返回最少的操作步骤将 x 计算到 y
// example
// Input X = 2, Y = 3, Output: 2

// 思路(贪心算法)
// 1. X > Y: 只能一步一步 -1得到
// 2. X == Y: 不用操作 return 0
// 3. X < Y:(看Y是奇数还是偶数,偶数直接/2，奇数+1后除2)
// 为什么是Y/2而不是X * 2呢，因为倒过来会发生的情况比较少
// 如果是X进行*2或者-1运算运算，那么在X为任何值的时候都可以进行两种操作
// 而倒过来计算，Y进行除2或者+1运算，那么当Y为奇数的时候，就不能进行除2运算。这样就减少了一些可能性。

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
