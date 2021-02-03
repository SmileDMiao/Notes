## context
---
> context是 goroutine 的上下文，包含 goroutine 的运行状态、环境、现场等信息。context 主要用来在 goroutine 之间传递上下文信息，包括: 取消信号, 超时时间, 截止时间, k-v 等。

1. 不要将 Context 塞到结构体里。直接将 Context 类型作为函数的第一参数，而且一般都命名为 ctx。
2. 不要向函数传入一个 nil 的 context，如果你实在不知道传什么，标准库给你准备好了一个 context: todo。
3. 不要把本应该作为函数参数的类型塞到 context 中，context 存储的应该是一些共同的数据。例如: 登陆的 session cookie 等。
4. 同一个 context 可能会被传递到多个 goroutine，别担心，context 是并发安全的。

## API
---
`Context 是一个接口，定义了 4 个方法，它们都是幂等的。也就是说连续多次调用同一个方法，得到的结果都是相同的`
```go
func Background() Context

func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
```

```go
type Context interface {
	// 当 context 被取消或者到了 deadline，返回一个被关闭的 channel
	Done() <-chan struct{}

	// 在 channel Done 关闭后，返回 context 取消原因
	Err() error

	// 返回 context 是否会被取消以及自动取消时间（即 deadline）
	Deadline() (deadline time.Time, ok bool)

	// 获取 key 对应的 value
	Value(key interface{}) interface{}
}
```

1. Done() 返回一个 channel，可以表示 context 被取消的信号: 当这个 channel 被关闭时，说明 context 被取消了。这是一个只读的channel。读一个关闭的 channel 会读出相应类型的零值。这是一个 receive-only 的 channel。因此在子协程里读这个 channel，除非被关闭，否则读不出来任何东西，子协程从 channel 里读出了值（零值）后，就可以做一些收尾工作，尽快退出。
2. Err() 返回一个错误，表示 channel 被关闭的原因。例如是被取消，还是超时。
3. Deadline() 返回 context 的截止时间，通过此时间，函数就可以决定是否进行接下来的操作，如果时间太短，就可以不往下做了，否则浪费系统资源。也可以用这个 deadline 来设置一个 I/O 操作的超时时间。
4. Value() 获取之前设置的 key 对应的 value。


`Background()和TODO()`
1. Background()和TODO()，这两个函数分别返回一个实现了Context接口的background和todo。
2. Background()主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。
3. TODO()，如果我们不知道该使用什么Context的时候，可以使用这个。
4. background和todo本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context。

`WithCancel`
WithCancel返回带有新Done通道的父节点的副本。当调用返回的cancel函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论先发生什么情况。取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。

`WithDeadline`
返回父上下文的副本，并将deadline调整为不迟于d。如果父上下文的deadline已经早于d，则WithDeadline(parent, d)在语义上等同于父上下文。当截止日过期时，当调用返回的cancel函数时，或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。

`WithTimeout`
WithTimeout返回`WithDeadline(parent, time.Now().Add(timeout))`。
取消此上下文将释放与其相关的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel，通常用于数据库或者网络连接的超时控制

## 作用
---
1. 传递共享数据(withValue)
2. 取消goroutine(cancel())
3. 防止goroutine泄漏(cancel())