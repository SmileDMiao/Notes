## Map底层实现
---
map 的设计也被称为 "The dictionary problem", 它的任务是设计一种数据结构用来维护一个集合的数据, 并且可以同时对集合进行增删查改的操作。最主要的数据结构有两种:
1. 哈希查找表(Hash Table)
2. 搜索树(Search Tree)

`Hash Table`
哈希查找表用一个哈希函数将 key 分配到不同的桶(bucket, 也就是数组的不同 index)。这样, 开销主要在哈希函数的计算以及数组的常数访问时间。在很多场景下, 哈希查找表的性能很高。哈希查找表一般会存在 "碰撞" 的问题, 就是说不同的 key 被哈希到了同一个 bucket。一般有两种应对方法: 链表法和开放地址法。链表法将一个 bucket 实现成一个链表, 落在同一个 bucket 中的 key 都会插入这个链表。开放地址法则是碰撞发生后, 通过一定的规律，在数组的后面挑选 "空位", 用来放置新的 key。

`Search Tree`
搜索树法一般采用自平衡搜索树, 包括: AVL 树, 红黑树。

## 结构
---
`Go 语言采用的是哈希查找表, 并且使用链表解决哈希冲突`
```go
type hmap struct {
  // 元素个数, 调用 len(map) 时直接返回此值
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	// buckets 的对数 log_2
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	// overflow 的 bucket 近似数
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	// 计算 key 的哈希的时候会传入哈希函数
	hash0     uint32 // hash seed
  
  // 指向 buckets 数组, 大小为 2^B
  // 如果元素个数为0, 就为 nil
	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	// 扩容的时候, buckets 长度会是 oldbuckets 的两倍
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	
	// 指示扩容进度, 小于此地址的 buckets 迁移完成
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}
```

`buckets 是一个指针, 最终它指向的是一个结构体`
```go
type bmap struct {
	tophash [bucketCnt]uint8
}
```

`但这只是表面的结构, 编译期间会动态地创建一个新的结构`
```go
type bmap struct {
	topbits  [8]uint8
	keys     [8]keytype
	values   [8]valuetype
	pad      uintptr
	overflow uintptr
}
```

`mapextra`
```go
type mapextra struct {
	// overflow[0] contains overflow buckets for hmap.buckets.
	// overflow[1] contains overflow buckets for hmap.oldbuckets.
	overflow [2]*[]*bmap

	// nextOverflow 包含空闲的 overflow bucket，这是预分配的 bucket
	nextOverflow *bmap
}
```
![IMAGE](resources/17BDBED5046833BCA2676AC1C6517D2A.jpg =744x548)


## key定位
---
1. map中的数据被存放于一个数组中的, 数组的元素是桶(bucket), 每个桶至多包含8个键值对数据。哈希值低位(low-order bits)用于选择桶, 哈希值高位(high-order bits)用于在一个独立的桶中区别出键。
2. 首先计算出待查找 key 的哈希, 使用低 5 位找到对应的bucket, 使用高 8 位 对应十进制在bucket 中寻找 tophash 值(HOB hash)与之匹配的 key。
3. 如果在 bucket 中没找到, 并且 overflow 不为空, 还要继续去 overflow bucket 中寻找, 直到找到或是所有的 key 槽位都找遍了, 包括所有的 overflow bucket


## 扩容
---
map 扩容的时机: 在向 map 插入新 key 的时候, 会进行条件检测, 符合下面这 2 个条件, 就会触发扩容:

1. 装载因子超过阈值, 源码里定义的阈值是 6.5。
2. overflow 的 bucket 数量过多: 当 B 小于 15, 也就是 bucket 总数 2^B 小于 2^15 时, 如果 overflow 的 bucket 数量超过 2^B; 当 B >= 15，也就是 bucket 总数 2^B 大于等于 2^15, 如果 overflow 的 bucket 数量超过 2^15。


`解释`
第 1 点: 我们知道，每个 bucket 有 8 个空位，在没有溢出，且所有的桶都装满了的情况下，装载因子算出来的结果是 8。因此当装载因子超过 6.5 时，表明很多 bucket 都快要装满了，查找效率和插入效率都变低了。在这个时候进行扩容是有必要的。

第 2 点: 是对第 1 点的补充，就是说在装载因子比较小的情况下，这时候 map 的查找和插入效率也很低，而第 1 点识别不出来这种情况。表面现象就是计算装载因子的分子比较小，即 map 里元素总数少，但是 bucket 数量多(真实分配的 bucket 数量多，包括大量的 overflow bucket)。

不难想像造成这种情况的原因: 不停地插入、删除元素。先插入很多元素，导致创建了很多 bucket，但是装载因子达不到第 1 点的临界值，未触发扩容来缓解这种情况。之后，删除元素降低元素总数量，再插入很多元素，导致创建很多的 overflow bucket，但就是不会触犯第 1 点的规定，怎么办？overflow bucket 数量太多，导致 key 会很分散，查找插入效率低得吓人，因此出台第 2 点规定。这就像是一座空城，房子很多，但是住户很少，都分散了，找起人来很困难。

`对于命中条件 1，2 的限制，都会发生扩容。但是扩容的策略并不相同，毕竟两种条件应对的场景不同`
1. 对于条件 1，元素太多，而 bucket 数量太少，很简单：将 B 加 1，bucket 最大数量（2^B）直接变成原来 bucket 数量的 2 倍。于是，就有新老 bucket 了。注意，这时候元素都在老 bucket 里，还没迁移到新的 bucket 来。而且，新 bucket 只是最大数量变为原来最大数量（2^B）的 2 倍（2^B * 2）。

2. 对于条件 2，其实元素没那么多，但是 overflow bucket 数特别多，说明很多 bucket 都没装满。解决办法就是开辟一个新 bucket 空间，将老 bucket 中的元素移动到新 bucket，使得同一个 bucket 中的 key 排列地更紧密。这样，原来，在 overflow bucket 中的 key 可以移动到 bucket 中来。结果是节省空间，提高 bucket 利用率，map 的查找和插入效率自然就会提升。


```go
// 触发扩容
if !h.growing() && (overLoadFactor(h.count+1, h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
	hashGrow(t, h)
	goto again // Growing the table invalidates everything, so try again
}

// 装载因子过大
// overLoadFactor reports whether count items placed in 1<<B buckets is over loadFactor.
func overLoadFactor(count int, B uint8) bool {
	return count > bucketCnt && uintptr(count) > loadFactorNum*(bucketShift(B)/loadFactorDen)
}

// overflow bucket过多
func tooManyOverflowBuckets(noverflow uint16, B uint8) bool {
	if B > 15 {
		B = 15
	}
	// The compiler doesn't see here that B < 16; mask B to generate shorter shift code.
	return noverflow >= uint16(1)<<(B&15)
}
```

1. 由于 map 扩容需要将原有的 key/value 重新搬迁到新的内存地址，如果有大量的 key/value 需要搬迁，会非常影响性能。因此 Go map 的扩容采取了一种称为“渐进式”地方式，原有的 key 并不会一次性搬迁完毕，每次最多只会搬迁 2 个 bucket。

2. 搬迁的目的就是将老的 buckets 搬迁到新的 buckets。而通过前面的说明我们知道，应对条件 1，新的 buckets 数量是之前的一倍，应对条件 2，新的 buckets 数量和之前相等。
对于条件 2，从老的 buckets 搬迁到新的 buckets，由于 bucktes 数量不变，因此可以按序号来搬，比如原来在 0 号 bucktes，到新的地方后，仍然放在 0 号 buckets。

3. 对于条件 1，就没这么简单了。要重新计算 key 的哈希，才能决定它到底落在哪个 bucket。例如，原来 B = 5，计算出 key 的哈希后，只用看它的低 5 位，就能决定它落在哪个 bucket。扩容后，B 变成了 6，因此需要多看一位，它的低 6 位决定 key 落在哪个 bucket。这称为 rehash。因此，某个 key 在搬迁前后 bucket 序号可能和原来相等，也可能是相比原来加上 2^B（原来的 B 值），取决于 hash 值 第 6 bit 位是 0 还是 1。

4. map 在扩容后，会发生 key 的搬迁，原来落在同一个 bucket 中的 key，搬迁后，有些 key 就要远走高飞了（bucket 序号加上了 2^B）。而遍历的过程，就是按顺序遍历 bucket，同时按顺序遍历 bucket 中的 key。搬迁后，key 的位置发生了重大的变化，有些 key 飞上高枝，有些 key 则原地不动。这样，遍历 map 的结果就不可能按原来的顺序了。

## map元素不可取址
---
```go
// 编译失败
func main() {
  m := make(map[string]int)

  fmt.Println(&m["qcrao"])
}
```
如果通过其他 hack 的方式, 例如 unsafe.Pointer 等获取到了 key 或 value 的地址, 也不能长期持有, 因为一旦发生扩容, key 和 value 的位置就会改变, 之前保存的地址也就失效了

## 线程安全
---
> map 并不是一个线程安全的数据结构, 同时读写一个 map 是未定义的行为, 如果被检测到, 会直接 panic。在查找 赋值 遍历 删除的过程中都会检测写标志, 一旦发现写标志置位等于1, 则直接 panic。赋值和删除函数在检测完写标志是复位之后, 先将写标志位置位, 才会进行之后的操作

```go
// 检测写标志
if h.flags&hashWriting != 0 {
	throw("concurrent map writes")
}

// 复位
h.flags ^= hashWriting
```

sync.RWMutex
1. 读之前调用 RLock() 函数, 读完之后调用 RUnlock() 函数解锁
2. 写之前调用 Lock() 函数, 写完之后, 调用 Unlock() 解锁

```go
type SafeMap struct {
	sync.RWMutex
	Map map[int]int
}

func main() {
	safeMap := newSafeMap(10)

	for i := 0; i < 100000; i++ {
		go safeMap.writeMap(i, i)
		go safeMap.readMap(i)
	}

}

func newSafeMap(size int) *SafeMap {
	sm := new(SafeMap)
	sm.Map = make(map[int]int)
	return sm

}

func (sm *SafeMap) readMap(key int) int {
	sm.RLock()
	value := sm.Map[key]
	sm.RUnlock()
	return value
}

func (sm *SafeMap) writeMap(key int, value int) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}
```

sync.Map
```go
func main() {
	// scene := sync.Map{}
	var scene sync.Map

	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	fmt.Println(scene.Load("london"))

	scene.Delete("london")

	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
```

## 比较两个map
---
map 深度相等的条件:
1. 都为 nil
2. 非空 长度相等 指向同一个 map 实体对象
3. 相应的 key 指向的 value "深度" 相等
4. 直接将使用 map1 == map2 是错误的, 这种写法只能比较 map 是否为 nil。

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(reflect.DeepEqual(m1, m2))
}
```

## map实现两种get操作
---
map有两种取值操作, 带 comma 和 不带 comma, 当要查询的 key 是否在 map 里, 带 comma 的用法会返回一个 bool 型变量提示 key 是否在 map 中; 而不带 comma 的语句则会返回一个 key 类型的零值。如果 key 是 int 型就会返回 0, 如果 key 是 string 类型, 就会返回空字符串。
```go
// src/runtime/map.go
// 两个一样的方法, 名字不一样
// 根据 key 的不同类型, 编译器还会将查找 插入 删除的函数用更具体的函数替换, 以优化效率
func mapaccess1(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer
func mapaccess2(t *maptype, h *hmap, key unsafe.Pointer) (unsafe.Pointer, bool) 
```