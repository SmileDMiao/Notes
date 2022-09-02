### strings
---
```go
// 替换字符串
strings.Replace(str, "6", "9", 1)

// string返回包含字符串的index -1:不存在
strings.Index(S, B)

// slice join into string 
word := strings.Join(words[:], "")

// 反转string
number_rune := []rune(positive_string)
for i, j := 0, len(number_rune)-1; i < j; i, j = i+1, j-1 {
	number_rune[i], number_rune[j] = number_rune[j], number_rune[i]
}
reverse_string := string(number_rune)
```

### strconv
---
```go
// string转换为int
result, _ := strconv.Atoi(str)
// int转换为string
str := strconv.Itoa(num)
```

### slice
---
1. 原切片数据量大，在原切片基础上只操作一小部分，可以使用 copy 替代
```go
// 清空slice
s := []string{"a", "b", "c"}
s = s[:0]
s = nil

// slice sort
sort.Strings(words)
sort.Ints(nums)

// 判断两个slice是否相等
reflect.DeepEqual(s, words)

// 删除元素
a = append(a[:i], a[i+1:]...)
// 删除尾部元素
a = a[:len(a)-1]

// 插入元素
a := []int{1, 2, 3, 4}
a = append(a[:2], append([]int{5}, a[2:]...)...)
// 头部插入
a = append([]int{5}, a...)
```

### reflect
---
```go
// DeepEqual: 判断是否深度相等
reflect.DeepEqual(s, words)
```