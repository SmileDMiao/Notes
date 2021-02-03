> slice的底层数据是数组, slice是对数组的封装，它描述一个数组的片段。两者都可以通过下标来访问单个元素。数组是定长的, 长度定义好之后, 不能再更改。在 Go 中，数组是不常见的, 因为其长度是类型的一部分, 限制了它的表达能力。而切片则非常灵活, 它可以动态地扩容, 切片的类型和长度无关。

1. 数组就是一片连续的内存, slice 实际上是一个结构体, 包含三个字段: 长度 容量 底层数组
2. 底层数组是可以被多个 slice 同时指向的, 因此对一个 slice 的元素进行操作是有可能影响到其他 slice 的


```go
// runtime/slice.go
type slice struct {
  array unsafe.Pointer // 元素指针
  len   int // 长度 
  cap   int // 容量
}
```

```go
func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
	fmt.Println(slice[:2:6])
}

// [2 3 20]
// [4 5 6 7 100 200]
// [0 1 2 3 20 5 6 7 100 9]
```

## 切片的扩容
> 一般都是在向 slice 追加了元素之后, 才会引起扩容。追加元素调用的是 append 函数。使用 append 可以向 slice 追加元素, 实际上是往底层数组添加元素。但是底层数组的长度是固定的, 如果索引 len-1 所指向的元素已经是底层数组的最后一个元素, 就没法再添加了。这时, slice 会迁移到新的内存位置, 新底层数组的长度也会增加, 这样就可以放置新增的元素。同时，为了应对未来可能再次发生的 append 操作, 新的底层数组的长度, 也就是新 slice 的容量是留了一定的 buffer 的。否则，每次添加元素的时候, 都会发生迁移, 成本太高。


这里要注意的是, append函数执行完后, 返回的是一个全新的 slice, 并且对传入的 slice 并不影响。

测试例子
```go
// 查看slice的扩容
// 1 2 4 8 16 32 64 128 256 512 1024 1280 1696 2304
func main() {
	s := make([]int, 0)

	oldCap := cap(s)

	for i := 0; i < 2048; i++ {
		s = append(s, i)

		newCap := cap(s)

		if newCap != oldCap {
			fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i-1, oldCap, i, newCap)
			oldCap = newCap
		}
	}
}

// len: 5 cap:6
func main() {
	s := []int{1, 2}
	s = append(s, 4, 5, 6)
	fmt.Printf("len=%d, cap=%d", len(s), cap(s))
}

// [5 7 9] [5 7 9 12] [5 7 9 12]
func main() {
	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)
	x := append(s, 11)
	y := append(s, 12)
	fmt.Println(s, x, y)
}
```




和扩容相关的源码部分
```go
// 参数依次是: 元素的类型, 老的 slice, 新 slice 最小求的容量
func growslice(et *_type, old slice, cap int) slice {
  ...
	newcap := old.cap
	doublecap := newcap + newcap
	if cap > doublecap {
		newcap = cap
	} else {
		if old.len < 1024 {
			newcap = doublecap
		} else {
			// Check 0 < newcap to detect overflow
			// and prevent an infinite loop.
			for 0 < newcap && newcap < cap {
				newcap += newcap / 4
			}
			// Set newcap to the requested cap when
			// the newcap calculation overflowed.
			if newcap <= 0 {
				newcap = cap
			}
		}
	}
}
...

// 分配内存
func roundupsize(size uintptr) uintptr {
	if size < _MaxSmallSize {
		if size <= smallSizeMax-8 {
			return uintptr(class_to_size[size_to_class8[divRoundUp(size, smallSizeDiv)]])
		} else {
			return uintptr(class_to_size[size_to_class128[divRoundUp(size-smallSizeMax, largeSizeDiv)]])
		}
	}
	if size+_PageSize < size {
		return size
	}
	return alignUp(size, _PageSize)
}
```



append的时候发生扩容的动作:
1. append单个元素, 或者append少量的多个元素, 这里的少量指double之后的容量能容纳, 这样就会走以下扩容流程, 不足1024, 双倍扩容, 超过1024的, 1.25倍扩容
2. 若是append多个元素, 且double后的容量不能容纳，直接使用预估的容量
3. 此外，以上两个分支得到新容量后, 均需要根据slice的类型size, 算出新的容量所需的内存情况capmem, 然后再进行capmem向上取整, 得到新的所需内存, 除上类型size, 得到真正的最终容量, 作为新的slice的容量