### 347. Top K Frequent Elements
> question

找出数组中出现最频繁的k个元素

> example

Input: nums = [1,1,1,2,2,3], k = 2 Output: [1,2]

> 思路(map+slice)

map保存数字与之对应的数量，数组保存map所有的key，然后根据map的value对数组排序, 获取数组的前k个元素
