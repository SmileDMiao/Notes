// 227. Basic Calculator II
// question
// 计算字符串表达式(+-*/)
// example
// Input: "3+2*2"; Output 7

// TODO

package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func calculate(s string) int {
	ans, cur := 0, 0
	stack := []int{}
	op := '+'
	for i, ch := range s {
		if ch-'0' >= 0 && ch-'0' <= 9 {
			cur = cur*10 + int(ch-'0')
		}
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' || i == len(s)-1 {
			if op == '+' {
				stack = append(stack, cur)
			} else if op == '-' {
				stack = append(stack, -cur)
			} else if op == '*' {
				t := cur * stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, t)
			} else if op == '/' {
				t := stack[len(stack)-1] / cur
				stack = stack[:len(stack)-1]
				stack = append(stack, t)
			}
			op = rune(ch)
			cur = 0
		}
	}

	for _, n := range stack {
		ans += n
	}
	return ans
}

func calculate(s string) int {
	operator := []string{}
	numstring := regexp.MustCompile("[\\+\\-\\*\\/\\s]+").Split(s, -1)

	numsint := make([]int, 0)
	for _, v := range numstring {
		vint, error := strconv.Atoi(v)
		if error != nil {
		} else {
			numsint = append(numsint, vint)
		}
	}

	for _, ch := range s {
		o := string(ch)
		if o == "+" || o == "-" || o == "*" || o == "/" {
			operator = append(operator, o)
		}
	}

	special := make([][]int, 0)
	var left, right int
	i := 0
	var start bool = false

	for i < len(operator) {
		if operator[i] == "*" {
			if start == false {
				start = true
				left = i
			}
		}
		if operator[i] == "/" {
			if start == false {
				start = true
				left = i
			}
		}
		if operator[i] == "+" || operator[i] == "-" {
			if start == true {
				start = false
				right = i
				special = append(special, []int{left, right})
			}
		}
		if i == len(operator)-1 {
			if start == true {
				right = i + 1
				special = append(special, []int{left, right})
			}
		}
		i++
	}

	kk := make([]int, 0)
	for i := 0; i < len(special); i++ {
		tmp := calculatePart(special[i], operator, numsint)
		kk = append(kk, tmp)
	}

	for i := 0; i < len(special); i++ {
		array := numsint[special[i][0]:(special[i][len(special[i])-1] + 1)]
		for j := 0; j < len(array); j++ {
			if j == len(array)-1 {
				array[j] = kk[i]
			} else {
				array[j] = -1
			}
		}
	}

	for i := 0; i < len(operator); {
		if operator[i] == "/" || operator[i] == "*" {
			operator = append(operator[:i], operator[i+1:]...)
		} else {
			i++
		}
	}

	for i := 0; i < len(numsint); {
		if numsint[i] == -1 {
			numsint = append(numsint[:i], numsint[i+1:]...)
		} else {
			i++
		}
	}

	fmt.Println(numsint)
	fmt.Println(kk)
	fmt.Println(special)

	result := numsint[0]
	for i := 1; i < len(numsint); i++ {
		if operator[i-1] == "+" {
			result += numsint[i]
		}
		if operator[i-1] == "-" {
			result -= numsint[i]
		}
	}

	return result
}

func calculatePart(nums []int, operator []string, numsint []int) int {
	var op []string
	var array []int

	array = numsint[nums[0]:(nums[len(nums)-1] + 1)]
	op = operator[nums[0]:nums[len(nums)-1]]

	tmp := array[0]

	for i := 1; i < len(array); i++ {
		if op[i-1] == "*" {
			tmp *= array[i]
		}
		if op[i-1] == "/" {
			tmp /= array[i]
		}
	}
	return tmp
}
