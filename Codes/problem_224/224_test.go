package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem224(t *testing.T) {
	s := "(1+(4+5+2)-3)+(6+8)"

	fmt.Println(calculate(s))
}
