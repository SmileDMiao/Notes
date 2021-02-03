## 接口
---
接口是一个自定义类型, 它是一组方法的集合, 要有方法为接口类型就被认为是该接口。

1. 接口本质是一种自定义类型
2. 接口是一种特殊的自定义类型, 其中没有数据成员，只有方法(也可以为空)


接口是完全抽象的, 因此不能将其实例化。然而, 可以创建一个其类型为接口的变量, 它可以被赋值为任何满足该接口类型的实际类型的值。接口的重要特性是:

1. 只要某个类型实现了接口所有的方法，那么我们就说该类型实现了此接口。该类型的值可以赋给该接口的值。
2. 作为1的推论, 任何类型的值都可以赋值给空接口interface{}。

## 空接口与nil
---
1. 空接口(interface{})不包含任何的method, 正因为如此, 所有的类型都实现了interface{}, interface{}可以存储任意类型的数值
2. nil只能赋值给指针 channel func interface map slice类型的变量。如果未遵循这个规则, 则会引发panic。
3. `(*interface{})(nil)`是将nil转成interface类型的指针, 其实得到的结果仅仅是空接口类型指针并且它指向无效的地址, 也就是空接口类型指针而不是空指针

```go
func main() {
	var val interface{} = (*interface{})(nil)
	if val == nil {
		fmt.Println("val is nil")
	} else {
		fmt.Println("val is not nil")
	}
}
```


## 接口变量存储的类型
---
```go
// `value, ok = element.(T)` value是变量的值, ok是一个bool类型, element是interface变量, T是断言的类型。如果element里面确实存储了T类型的数值, 那么ok返回true, 否则返回false。
func main() {
	i := 10
	var a interface{} = i
	s, ok := a.(string)

	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(ok)
}
```

## 接口实现
---
用来说明是否我们一个类型的值或者指针实现了该接口的规则:

1. 类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集(接收者是指针 *T 时，接口的实例必须是指针)
2. 类型 T 的可调用方法集包含接受者为 T 的所有方法(接收者是值 T 时, 接口的实例可以是指针也可以是值)
3. 类型 T 的可调用方法集不包含接受者为 *T 的方法

### 接收者的方法
---
Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。
与普通函数不同, 接收者为指针类型和值类型的方法, 指针类型和值类型的变量均可相互调用  
方法的定义格式如下:
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}

什么时候应该使用指针类型接收者:
1. 需要修改接收者中的值
2. 接收者是拷贝代价比较大的大对象
3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

## iface 和 eface
---
`iface` 和 `eface` 都是 Go 中描述接口的底层结构体, 区别在于 iface 描述的接口包含方法, 而 `face` 则是不包含任何方法的空接口: `interface{}`

#### iface
iface 内部维护两个指针
1. tab 指向一个 itab 实体, 它表示接口的类型以及赋给这个接口的实体类型
2. data 则指向接口具体的值, 一般而言是一个指向堆内存的指针

#### itab
1. _type: 字段描述了实体的类型, 包括内存对齐方式, 大小等
2. inter 字段则描述了接口的类型
3. fun 字段放置和接口方法对应的具体数据类型的方法地址, 实现接口调用方法的动态分派
4. 为什么 fun 数组的大小为 1, 要是接口定义了多个方法可怎么办？
这里存储的是第一个方法的函数指针, 如有更多的方法, 在它之后的内存空间里继续存储。从汇编角度来看, 通过增加地址就能获取到这些函数指针, 没什么影响

#### interfacetype
1. _type: 描述 Go 语言中各种数据类型的结构体, Go 语言各种数据类型都是在 _type 字段的基础上, 增加一些额外的字段来进行管理的：
2. mhdr: 表示接口所定义的函数列表
3. pkgpath: 记录定义了接口的包名

#### eface
1. _type: 空接口所承载的具体的实体类型
2. data: 描述了具体的值

```go
// tab: 接口类型以及赋给这个接口的实体类型
// data: 向接口具体的值
type iface struct {
  tab  *itab 
  data unsafe.Pointer
}

// _type 字段描述了实体的类型，包括内存对齐方式，大小等
// inter 字段则描述了接口的类型。
// fun 字段放置和接口方法对应的具体数据类型的方法地址, 实现接口调用方法的动态分派
type itab struct {
  inter *interfacetype
  _type *_type
  hash  uint32 // copy of _type.hash. Used for type switches.
  _     [4]byte
  fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}

//  _type: 描述 Go 语言中各种数据类型的结构体
// mhdr: 表示接口所定义的函数列表
// pkgpath: 记录定义了接口的包名
type interfacetype struct {
  typ     _type
  pkgpath name
  mhdr    []imethod
}

// _type: 空接口所承载的具体的实体类型
// data: 描述了具体的值
type eface struct {
  _type *_type
  data  unsafe.Pointer
}

// size: 类型大小
// hash: 类型的hash值
// tflag: 类型的 flag, 和反射相关
// align: 内存对齐相关
type _type struct {
  size       uintptr
  ptrdata    uintptr // size of memory prefix holding all pointers
  hash       uint32
  tflag      tflag
  align      uint8
  fieldAlign uint8
  kind       uint8
  // function for comparing objects of this type
  // (ptr to object A, ptr to object B) -> ==?
  equal func(unsafe.Pointer, unsafe.Pointer) bool
  // gcdata stores the GC type data for the garbage collector.
  // If the KindGCProg bit is set in kind, gcdata is a GC program.
  // Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
  gcdata    *byte
  str       nameOff
  ptrToThis typeOff
}
```

## 值接收者和指针接收者
---
1. 给一个函数添加一个接收者, 那么它就变成了方法。接收者可以是值接收者, 也可以是指针接收者。
2. 在调用方法的时候, 值类型既可以调用值接收者的方法, 也可以调用指针接收者的方法
3. 指针类型既可以调用指针接收者的方法, 也可以调用值接收者的方法。

当类型和方法的接收者类型不同时, 其实是编译器在背后做了一些工作:
| | 值类型调用者 | 指针接收者 | 
| :-----: | :----: | :----: |
| 值类型调用者 | 方法会使用调用者的一个副本, 类似于 "传值" | 使用值的引用来调用方法, qcrao.growUp() 实际上是 (&qcrao).growUp() |
| 指针类型调用者 | 指针被解引用为值, stefno.howOld() 实际上是 (*stefno).howOld() | 实际上也是 "传值", 方法里的操作会影响到调用者, 类似于指针传参, 拷贝了一份指针 |


1. 实现了接收者是值类型的方法, 相当于自动实现了接收者是指针类型的方法
2. 实现了接收者是指针类型的方法, 不会自动生成对应接收者是值类型的方法

```go
type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

func main() {
	var c coder = &Gopher{"Go"}
	c.code()
	c.debug()
}
```