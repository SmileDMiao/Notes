## CSP模型
---
与主流语言通过共享内存来进行并发控制方式不同, Go 语言采用了 CSP 模式。这是一种用于描述两个独立的并发实体通过共享的通讯 Channel（管道）进行通信的并发模型。
Golang 就是借用CSP模型的一些概念为之实现并发进行理论支持, 其实从实际上出发, go语言并没有, 完全实现了CSP模型的所有理论, 仅仅是借用了 process和channel这两个概念。process是在go语言上的表现就是 goroutine 是实际并发执行的实体, 每个实体之间是通过channel通讯来实现数据共享。

Go语言的CSP模型是由协程Goroutine与通道Channel实现：

1. Go协程goroutine: 是一种轻量线程, 它不是操作系统的线程, 而是将一个操作系统线程分段使用, 通过调度器实现协作式调度。是一种绿色线程, 微线程，它与Coroutine协程也有区别, 能够在发现堵塞后启动新的微线程。
2. 通道channel: 类似Unix的Pipe，用于协程之间通讯和同步。协程之间虽然解耦，但是它们和Channel有着耦合。


## 不要以共享内存的方式去通信而要以通信的方式去共享内存: Channel
---
要想解决临界资源安全的问题, 很多编程语言的解决方案都是同步。通过上锁的方式，某一时间段, 只能允许一个goroutine来访问这个共享数据, 当前goroutine访问完毕, 解锁后, 其他的goroutine才能来访问。

### channel:
通道可以被认为是Goroutines通信的管道。类似于管道中的水从一端到另一端的流动, 数据可以从一端发送到另一端，通过通道接收。
当多个Goroutine想实现共享数据的时候, 虽然也提供了传统的同步机制, 但是Go语言强烈建议的是使用Channel通道来实现Goroutines之间的通信。

Go语言中, 要传递某个数据给另一个goroutine(协程),可以把这个数据封装成一个对象, 然后把这个对象的指针传入某个channel中, 另外一个goroutine从这个channel中读出这个指针, 并处理其指向的内存对象。Go从语言层面保证同一个时间只有一个goroutine能够访问channel里面的数据。channel是引用类型的数据, 在作为参数传递的时候, 传递的是内存地址。

1. 用于goroutine, 传递消息
2. 每个channel都有相关联的数据类型, nil chan，不能使用
3. 本身channel就是同步的, 意味着同一时间, 只能有一条goroutine来操作
4. 阻塞: 发送数据(chan <- data)是阻塞的, 直到另一条goroutine, 读取数据来解除阻塞。读取数据(data <- chan)也是阻塞的。直到另一条goroutine, 写出数据解除阻塞。
5. 死锁: 如果Goroutine在一个通道上发送数据, 那么预计其他的Goroutine应该接收数据。如果这种情况不发生，那么程序将在运行时出现死锁。类似, 如果Goroutine正在等待从通道接收数据, 那么另一些Goroutine将会在该通道上写入数据, 否则程序将会死锁。

关闭后的通道有以下特点：
1. 对一个关闭的通道再发送值就会导致panic。
2. 对一个关闭的通道进行接收会一直获取值直到通道为空。
3. 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
4. 关闭一个已经关闭的通道会导致panic。

## Channel的应用
---
### 如何优雅的关闭channel
---
`关于 channel 的使用, 有几点不方便的地方`
1. 在不改变 channel 自身状态的情况下, 无法获知一个 channel 是否关闭
2. 关闭一个 closed channel 会导致 panic。如果关闭 channel 的一方在不知道 channel 是否处于关闭状态时就去贸然关闭 channel 是很危险的事情
3.向一个 closed channel 发送数据会导致 panic。如果向 channel 发送数据的一方不知道 channel 是否处于关闭状态时就去贸然向 channel 发送数据是很危险的事情

`关闭channel的规则`
1. 不要从一个 receiver 侧关闭 channel, 也不要在有多个 sender 时, 关闭 channel
2. 不要关闭或发送数据到已关闭的channel


`有两个不那么优雅地关闭 channel 的方法`
1. 使用 defer-recover 机制, 放心大胆地关闭 channel 或者向 channel 发送数据。即使发生了 panic, 有 defer-recover 在兜底
2. 使用 sync.Once 来保证只关闭一次

`具体情况`
1. 一个 sender 一个 receiver
2. 一个 sender  M 个 receiver
3. N 个 sender 一个 reciver
4. N 个 sender  M 个 receiver

对于 1和2, 只有一个 sender 的情况就不用说了, 直接从 sender 端关闭就好了, 没有问题

`N个sender一个receiver`
优雅关闭 channel 的方法是: the only receiver says "please stop sending more" by closing an additional signal channel
stopCh 就是信号 channel, 它本身只有一个 sender, 因此可以直接关闭它。senders 收到了关闭信号后, select 分支 `case <- stopCh` 被选中, 退出函数，不再发送数据。
其实并没有明确关闭 dataCh。一个 channel, 如果最终没有任何 goroutine 引用它, 不管 channel 有没有被关闭, 最终都会被 gc 回收

```go
func main() {
	const Max = 100000
	const NumSenders = 1000

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	// the receiver
	go func() {
		for value := range dataCh {
			if value == Max-1 {
				fmt.Println("send stop signal to senders.")
				close(stopCh)
				return
			}

			fmt.Println(value)
		}
	}()

	select {
	case <-time.After(time.Hour):
	}
}

```

`N个sender M个 receiver`
声明一个中间人的角色channel, 用它来接收 senders 和 receivers 发送过来的关闭 数据channel 请求。
这里将 中间人channel声明成了一个 缓冲型的 channel。假设声明的是一个非缓冲型的 channel, 那么第一个发送的关闭数据channel 请求可能会丢失。因为无论是 sender 还是 receiver 都是通过 select 语句来发送请求, 如果中间人所在的 goroutine 没有准备好, 那 select 语句就不会选中, 直接走 default 选项, 什么也不做。这样, 第一个关闭 dataCh 的请求就会丢失。

### 任务定时
---
```go
// time.After
// 一次性定时
select {
	case <-time.After(100 * time.Millisecond):
}

// time.Tick
// 循环定时
func worker() {
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-ticker:
			// 执行定时任务
			fmt.Println("执行 1s 定时任务")
		}
	}
}

```

### 解偶生产方消费方
---
```go
// 5 个工作协程在不断地从工作队列里取任务, 生产方只管往 channel 发送任务即可, 解耦生产方和消费方
func main() {
	taskCh := make(chan int, 100)
	go worker(taskCh)

	// 塞任务
	for i := 0; i < 10; i++ {
		taskCh <- i
	}

	// 等待 1 小时
	select {
	case <-time.After(time.Hour):
	}
}

func worker(taskCh <-chan int) {
	const N = 5
	// 启动 5 个工作协程
	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				task := <-taskCh
				fmt.Printf("finish task: %d by worker %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}
```

### 控制并发数
---
```go
// 缓冲通道
func main() {
	userCount := 10
	ch := make(chan bool, 2)
	for i := 0; i < userCount; i++ {
		ch <- true
		go Read(ch, i)
	}

	time.Sleep(time.Second)
}

func Read(ch chan bool, i int) {
	fmt.Printf("go func: %d\n", i)
	<-ch
}
```

## 深入Channel
---
### 字段
---
```go
type hchan struct {
	// chan 里元素数量
	qcount uint // total data in the queue
	
	// chan 底层循环数组的长度
	dataqsiz uint // size of the circular queue

	// 指向底层循环数组的指针
	// 只针对有缓冲的 channel
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	
	// chan中元素大小
	elemsize uint16
	
	// chan 是否被关闭的标志
	closed   uint32
	
	// chan 中元素类型
	elemtype *_type // element type
	
	// 已发送元素在循环数组中的索引
	sendx    uint   // send index
	
	// 已接收元素在循环数组中的索引
	recvx    uint   // receive index
	
	// 等待接收的 goroutine 队列
	recvq    waitq  // list of recv waiters
	
	// 等待发送的 goroutine 队列
	sendq    waitq  // list of send waiters

  // 保护 hchan 中所有字段
	lock mutex
}
```

1. buf 指向底层循环数组 只有缓冲型的 channel 才有
2. sendx recvx 均指向底层循环数组, 表示当前可以发送和接收的元素位置索引值(相对于底层数组)
3. sendq recvq 分别表示被阻塞的 goroutine, 这些 goroutine 由于尝试读取 channel 或向 channel 发送数据而被阻塞
4. waitq 是 sudog 的一个双向链表, 而 sudog 实际上是对 goroutine 的一个封装
    ```go
    type waitq struct {
        first *sudog
        last  *sudog
    }
    ```
5. lock 用来保证每个读 channel 或写 channel 的操作都是原子的

### Channel发送接收元素点本质
---
Channel里发送和接收操作本质上是值拷贝
```go
type user struct {
	name string
	age  int8
}

var u = user{name: "Ankur", age: 25}
var g = &u

func modifyUser(pu *user) {
	fmt.Println("modifyUser Received Vaule", pu)
	pu.name = "Anand"
}

func printUser(u <-chan *user) {
	time.Sleep(2 * time.Second)
	fmt.Println("printUser goRoutine called", <-u)
}

func main() {
	c := make(chan *user, 5)
	c <- g
	fmt.Println(g)
	// modify g
	g = &user{name: "Ankur Anand", age: 100}
	go printUser(c)
	go modifyUser(g)
	time.Sleep(5 * time.Second)
	fmt.Println(g)
}

// &{Ankur 25}
// modifyUser Received Vaule &{Ankur Anand 100}
// printUser goRoutine called &{Ankur 25}
// &{Anand 100}
```