#### strings
```go
// 替换字符串
strings.Replace(str, "6", "9", 1)

// string返回包含字符串的index -1:不存在
strings.Index(S, B)

// slice join into string 
word := strings.Join(words[:], "")
```

#### strconv
```go
// string转换为int
result, _ := strconv.Atoi(str)
// int转换为string
str := strconv.Itoa(num)
```

#### slice
```
// 清空slice
s := []string{"a", "b", "c"}
s = s[:0]

// slice sort
sort.Strings(words)

// 判断两个slice是否相等
reflect.DeepEqual(s, words)
```

#### reflect
```go
// DeepEqual: 判断是否深度相等
reflect.DeepEqual(s, words)
```