// 1323. Maximum 69 Number
// question
// 一个数组只包含6和9,只能将一个数字(6->9/9->6)改变，返回能得到的最大值
// example
// Input: nums = [1,3,5,6], target = 5, Output: 2

package main

import (
	"strconv"
	"strings"
)

func maximum69Number(num int) int {
	str := strconv.Itoa(num)
	str = strings.Replace(str, "6", "9", 1)
	result, _ := strconv.Atoi(str)
	return result
}
