### Go语言指针的限制
---
1. 不能进行数学运算
2. 不同类型的指针不能相互转换
3. 不同类型的指针不能使用 == 或 != 比较
4. 不同类型的指针变量不能相互赋值

### 非类型安全的指针unsafe
---
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

*uintptr*就是一个16进制的整数，这个数字表示对象的地址，但是uintptr没有指针的语义。
1. 如果一个对象只有一个 uintptr 表示的地址表示"引用"关系，那么这个对象会在GC时被无情的回收掉，那么uintptr表示一个野地址。
2. 如果uintptr表示的地址指向的对象发生了copy移动(比如协程栈增长，slice的扩容等)，那么uintptr也表示一个野地址。但是unsafe.Pointer 有指针语义，可以保护它所指向的对象在"有用"的时候不会被垃圾回收，并且在发生移动时候更新地址值。


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

### 字符串转byte数组会发生内存拷贝吗?
---
字符串转成切片，会产生拷贝。严格来说，只要是发生类型强转都会发生内存拷贝。
```go
package main

import (
 "fmt"
 "reflect"
 "unsafe"
)

func main() {
 a :="aaa"
 ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
 b := *(*[]byte)(unsafe.Pointer(&ssh))  
 fmt.Printf("%v",b)
}
```

```go
// StringHeader 是字符串在go的底层结构。
type StringHeader struct {
 Data uintptr
 Len  int
}

// SliceHeader 是切片在go的底层结构。

type SliceHeader struct {
 Data uintptr
 Len  int
 Cap  int
}
```
那么如果想要在底层转换二者，只需要把 StringHeader 的地址强转成 SliceHeader 就行。

1. unsafe.Pointer(&a)方法可以得到变量a的地址。
2. (*reflect.StringHeader)(unsafe.Pointer(&a)) 可以把字符串a转成底层结构的形式。
3. (*[]byte)(unsafe.Pointer(&ssh)) 可以把ssh底层结构体转成byte的切片的指针。再通过 *转为指针指向的实际内容。