// 347. Top K Frequent Elements
// question
// 找出数组中出现最频繁的k个元素
// example
// Input: nums = [1,1,1,2,2,3], k = 2 Output: [1,2]

// 思路1(map+slice)
// map保存数字与之对应的数量，数组保存map所有的key，然后根据map的value对数组排序

// TODO
package main

import (
	"container/heap"
	"fmt"
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

type KMaxHeap [][]int

func (h KMaxHeap) Len() int           { return len(h) }
func (h KMaxHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h KMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

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
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	val := heap.Pop(h)

	fmt.Println(val)

	// for i := 0; i < k; i++ {
	// 	val := heap.Pop(h)
	// 	// result = append(result, va)
	// 	fmt.Println(val)
	// }

	fmt.Println(result)
	return result
}

func main() {
	words := []int{-1, -1}
	topKFrequent1(words, 1)

	topKFrequent2(words, 1)

}
