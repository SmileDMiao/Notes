package leetcode

import (
	"container/heap"
	"sort"
)

func topKFrequent1(nums []int, k int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	var keys []int
	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return keys[0:k]
}

// 思路2(堆)
type KMaxHeap [][]int

func (h KMaxHeap) Len() int           { return len(h) }
func (h KMaxHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h KMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *KMaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *KMaxHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

func topKFrequent2(nums []int, k int) []int {
	frequency := map[int]int{}
	for _, num := range nums {
		frequency[num]++
	}

	h := &KMaxHeap{}
	result := []int{}
	for k, v := range frequency {
		heap.Push(h, []int{v, k})
	}

	for i := 0; i < k; i++ {
		val := heap.Pop(h).([]int)
		result = append(result, val[1])
	}

	return result
}
