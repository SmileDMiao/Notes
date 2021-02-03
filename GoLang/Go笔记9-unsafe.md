### Go语言指针的限制
1. 不能进行数学运算
2. 不同类型的指针不能相互转换
3. 不同类型的指针不能使用 == 或 != 比较
4. 不同类型的指针变量不能相互赋值


### 非类型安全的指针unsafe
1. 任何类型的指针都可以被转化为Pointer
2. Pointer可以被转化为任何类型的指针
3. uintptr可以被转化为Pointer
4. Pointer可以被转化为uintptr

指针:
1. *类型: 普通指针类型，用于传递对象地址，不能进行指针运算。
2. unsafe.Pointer: 通用指针类型，用于转换不同类型的指针，不能进行指针运算，不能读取内存存储的值(必须转换到某一类型的普通指针)。
3. uintptr: 用于指针运算，GC 不把 uintptr 当指针，uintptr 无法持有对象。uintptr 类型的目标会被回收。

```go
func Sizeof(x ArbitraryType) uintptr
func Offsetof(x ArbitraryType) uintptr
func Alignof(x ArbitraryType) uintptr
```

1. Sizeof 返回类型 x 所占据的字节数，但不包含 x 所指向的内容的大小。例如: 对于一个指针，函数返回的大小为 8 字节（64位机上），一个 slice 的大小则为 slice header 的大小。
2. Offsetof 返回结构体成员在内存中的位置离结构体起始处的字节数，所传参数必须是结构体的成员。
3. Alignof 返回 m，m 是指当类型进行内存对齐时，它分配到的内存地址能整除 m。


```go
func main() {
	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s)) // 9 9

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s)) // 20 20

	mp := make(map[string]int)
	mp["qcrao"] = 100
	mp["stefno"] = 18

	count := **(**int)(unsafe.Pointer(&mp))
	fmt.Println(count, len(mp)) // 2 2
}

```