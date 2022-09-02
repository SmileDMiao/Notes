## 执行顺序
---
多个 defer 出现的时候, 它是一个 "栈" 的关系, 也就是先进后出。一个函数中, 写在前面的 defer 会比写在后面的 defer 调用的晚。
```go
// C B A
func main() {
	defer func1()
	defer func2()
	defer func3()
}

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}
```

## 和 return 的顺序
---
先 return 后 defer
```go
func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	returnAndDefer()
}
```

## 有返回值的情况
---
```go
// 10: 只要声明函数的返回值变量名称, return之后执行defer, 依然可以修改值
func returnButDefer() (t int) { 
	defer func() {
		t = t * 10
	}()
	return 1
}

// 1
func returnDefer() int {
  t := 1
  defer func() {
    t += 3
  }()
  return t
}

func main() {
	fmt.Println(returnButDefer())
	fmt.Println(returnDefer())
}
```

## panic
---
```go
// 2 1 panic: 3
// 遇到 panic 时, 遍历 defer 并执行。在执行 defer 过程中: 遇到 recover 则停止 panic, 返回 recover 处继续往下执行。如果没有遇到 recover, 遍历 defer 后, 抛出 panic 信息。
func main() {
	defer_call()

	fmt.Println("main 正常结束")
}

func defer_call() {
	defer func() { fmt.Println("1") }()
	defer func() { fmt.Println("2") }()

	panic("3") 

	defer func() { fmt.Println("defer: panic 之后，永远执行不到") }()
}
```

```go
// 4 2 3 1
func main() {
	defer_call()

	fmt.Println("1")
}

func defer_call() {

	defer func() {
		fmt.Println("2")
		if err := recover(); err != nil {
			fmt.Println(3)
		}
	}()

	defer func() { fmt.Println("4") }()

	panic("5")

	defer func() { fmt.Println("6") }()
}
```

```go
// 1
func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(1)
		} else {
			fmt.Println("2")
		}
	}()

	defer func() {
		panic("3")
	}()

	panic("4")
}
```

## 函数参数包含子函数
---
```go
// 3 4 2 1
// 进入斩顺序, 第一个先进, 为了进入先求第二个参数, 同理第二个defer
func function(index int, value int) int {

	fmt.Println(index)

	return index
}

func main() {
	defer function(1, function(3, 0))
	defer function(2, function(4, 0))
}
```