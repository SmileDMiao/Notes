package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem1(t *testing.T) {
	first_string := "abcd"
	second_string := "cdabcdab"

	fmt.Println(repeatedStringMatch(first_string, second_string))
}
