package leetcode

import (
	"testing"
)

func Test_Problem1(t *testing.T) {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}

	topKFrequent1(words, 3)

	topKFrequent2(words, 3)
}
