### 692. Top K Frequent Words
> question

前K个出现最频繁的单词

> example

Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2 Output:  ["i", "love"]

> 思路1(map+slice)
map保存word与之对应的数量，数组保存所有words(去重了的)然后根据map的value对数组排序
