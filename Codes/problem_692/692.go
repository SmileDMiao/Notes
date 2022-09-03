package leetcode

import (
	"fmt"
	"sort"
)

func topKFrequent1(words []string, k int) []string {
	m := make(map[string]int)

	for _, v := range words {
		count := m[v] + 1
		m[v] = count
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if m[keys[i]] == m[keys[j]] {
			return keys[i] < keys[j]
		} else {
			return m[keys[i]] > m[keys[j]]
		}
	})

	fmt.Println(keys[0:k])
	return keys[0:k]
}

// 思路2(map+slice)
// map保存word与之对应的数量，数组是排序的，对应的index是对应的所有words出现次数的集合，反向遍历数组，挨个写入结果并计算次数
func topKFrequent2(words []string, k int) []string {
	if len(words) == 0 {
		return nil
	}
	var res []string
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}
	buckets := make([][]string, len(words)+1)
	for key, value := range m {
		buckets[value] = append(buckets[value], key)
	}

	fmt.Println(buckets)
	count := 0
	for i := len(words); i >= 0; i-- {
		b := buckets[i]
		if len(b) != 0 {
			sort.Strings(b)
			for _, s := range b {
				res = append(res, s)
				count++
				if count == k {
					return res
				}
			}
		}
	}
	return res
}
