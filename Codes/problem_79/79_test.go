package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem79(t *testing.T) {
	board := [][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}
	word := "ABCCED"

	fmt.Println(exist(board, word))
}
